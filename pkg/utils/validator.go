package utils

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
	"github.com/gin-gonic/gin"
	"github.com/stvp/rollbar"
	"github.com/go-playground/validator/v10"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)


// const list
const (
	ErrorTypeErrorValidation string = "validator.ValidationErrors"
)

// var list
var (
	ErrorInternalError = errors.New("whoops something went wrong")
	matchFirstCap      = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap        = regexp.MustCompile("([a-z0-9])([A-Z])")
)

// ToSnakeCase method change string to snakecase
func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// UcFirst method is for uppercase first letter of word
func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func IsInterfaceToJSON(i interface{}) bool {
	_, err := json.Marshal(i)
	if err != nil {
		return false
	}
	return true
}

func IsJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func IsJSONString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) == nil
}

// LcFirst method is to change lowercase word
func LcFirst(str string) string {
	return strings.ToLower(str)
}

// Split method is to split camelcase letter
func Split(src string) string {
	// don't split invalid utf8
	if !utf8.ValidString(src) {
		return src
	}
	var entries []string
	var runes [][]rune
	lastClass := 0
	class := 0
	// split into fields based on class of unicode character
	for _, r := range src {
		switch true {
		case unicode.IsLower(r):
			class = 1
		case unicode.IsUpper(r):
			class = 2
		case unicode.IsDigit(r):
			class = 3
		default:
			class = 4
		}
		if class == lastClass {
			runes[len(runes)-1] = append(runes[len(runes)-1], r)
		} else {
			runes = append(runes, []rune{r})
		}
		lastClass = class
	}

	for i := 0; i < len(runes)-1; i++ {
		if unicode.IsUpper(runes[i][0]) && unicode.IsLower(runes[i+1][0]) {
			runes[i+1] = append([]rune{runes[i][len(runes[i])-1]}, runes[i+1]...)
			runes[i] = runes[i][:len(runes[i])-1]
		}
	}
	// construct []string from results
	for _, s := range runes {
		if len(s) > 0 {
			entries = append(entries, string(s))
		}
	}

	for index, word := range entries {
		if index == 0 {
			entries[index] = UcFirst(word)
		} else {
			entries[index] = LcFirst(word)
		}
	}
	justString := strings.Join(entries, " ")
	return justString
}

// ExtraValidation model
type ExtraValidation struct {
	Tag     string
	Message string
}

// ValidationObject initialize default Validation object
var ValidationObject = []ExtraValidation{
	{Tag: "required", Message: "%s is required!"},
	{Tag: "max", Message: "%s cannot be more than %s!"},
	{Tag: "min", Message: "%s must be minimum %s!"},
	{Tag: "email", Message: "Invalid email format!"},
	{Tag: "len", Message: "%s must be %s characters long!"},
}

// MakeExtraValidation method is for registering new validator
func MakeExtraValidation(v []ExtraValidation) {
	for _, vObj := range v {
		ValidationObject = append(ValidationObject, vObj)
	}

}

// checkOccurance checks if param is involved in valdation message
func checkOccurance(msg string, word string, param string) (ans string) {
	reg := regexp.MustCompile("%s")

	matches := reg.FindAllStringIndex(msg, -1)
	if len(matches) == 2 {
		ans = fmt.Sprintf(msg, word, param)
	} else {
		ans = fmt.Sprintf(msg, word)
	}
	return
}

// ValidationErrorToText method changes FieldError to string
func ValidationErrorToText(e validator.FieldError) string {
	word := Split(e.Field())
	var result string
	for _, validate := range ValidationObject {
		if e.Tag() == validate.Tag {
			result = checkOccurance(validate.Message, word, e.Param())
		}
	}
	if result == "" {
		result = fmt.Sprintf("%s is not valid", word)
	}

	return result

}

// Errors method collects all errors and submits them to Rollbar
func Errors() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Only run if there are some errors to handle
		c.Next()
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				// Find out what type of error it is
				switch e.Type {
				case gin.ErrorTypePublic:
					if !c.Writer.Written() {
						c.JSON(c.Writer.Status(), gin.H{"Error": e.Error()})
					}
				case gin.ErrorTypeBind:
					errorType := reflect.TypeOf(e.Err).String()
					switch errorType {
					case ErrorTypeErrorValidation:
						errs := e.Err.(validator.ValidationErrors)
						list := make(map[string]string)

						for _, err := range errs {
							list[strings.ToLower(toSnakeCase(err.Field()))] = ValidationErrorToText(err)
						}
						status := http.StatusUnprocessableEntity
						c.JSON(status, gin.H{"error": list})
						c.Abort()
						return
					default:
						status := http.StatusUnprocessableEntity
						c.JSON(status, gin.H{"error": e.Err.Error()})
						c.Abort()
						return
					}

				default:
					// Log all other errors
					rollbar.RequestError(rollbar.ERR, c.Request, e.Err)
				}

			}
			// If there was no public or bind error, display default 500 message
			if !c.Writer.Written() {
				c.JSON(http.StatusInternalServerError, gin.H{"error": ErrorInternalError.Error()})
			}
		}
	}
}

func Validate(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		return err
	}

	return nil
}