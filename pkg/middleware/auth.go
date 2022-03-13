package middleware

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/iromli/go-itsdangerous"
	"github.com/EZChain-core/price-service/config"
	b64 "encoding/base64"
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"strings"
	"unicode/utf8"
	"time"
	jsoniter "github.com/json-iterator/go"
	"github.com/EZChain-core/price-service/logger"
	"github.com/EZChain-core/price-service/pkg/utils"


)

var (
	 json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func AuthenticationRequired(appConfig *config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context)	 {
		DecryptCookie(appConfig, c, "cookie", nil)
	}
}

func FeatureCheck(flag *utils.FeatureFlag) gin.HandlerFunc {
	return func(c *gin.Context)	 {
		// check if flag is enable
		if(flag != nil && flag.IsEnable(utils.FEATURE_USER_NAME, utils.FEATURE_USER_NAME)) == true {
			email, _ := c.Get("email")
			options := map[string]string{
				"userkey": email.(string),
			}


			if isFlag, err := flag.InquiryWithContext(utils.FEATURE_USER_NAME, options); isFlag == false || err != nil {
				c.JSON(http.StatusForbidden, gin.H{
					"error_code": http.StatusForbidden,
					"message":    "Forbiden, You cannot access this url",
					"success":    false,
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

func DecryptCookieRaw(appConfig *config.AppConfig, cookie string) *UserPayload {
	if appConfig.UseAuth == true {
		var userPayload *UserPayload

		signer := itsdangerous.NewTimestampSignature(
			appConfig.Secret,
			appConfig.Salt,
			appConfig.Sep,
			"hmac",
			nil,
			nil,
		)

		// unsign cookies to decrypt user info
		signed, err := signer.Unsign(cookie, 0)

		if err != nil {
			logger.Error(err)
			return nil
		}

		// append == into string to convert to base64
		base64String := signed[1:len(signed)]
		base64String = base64String + strings.Repeat("=", (4-(utf8.RuneCountInString(base64String)%4))%4)

		data, err := b64.URLEncoding.DecodeString(base64String)

		if err != nil {
			logger.Error(err)
			return nil
		}

		b := bytes.NewReader(data)
		read, err := zlib.NewReader(b)

		defer read.Close()

		if err != nil {
			logger.Error(err)
			return nil
		}

		payload, err := ioutil.ReadAll(read)
		if err != nil {
			logger.Error(err)
			return nil
		}
		logger.Debug(fmt.Sprint("payload decyprt is %s", string(payload)))

		if err := json.Unmarshal(payload, &userPayload); err != nil {
			logger.Error(err)
			return nil
		}
		// check expire of token
		expire := (*userPayload.Expires).Unix()

		if expire <= time.Now().Unix() {
			logger.Error(err)
			return nil
		}
		return userPayload
	}

	return nil
}


func DecryptCookie(appConfig *config.AppConfig, c *gin.Context, decryptType string, options map[string]string)  {
	if appConfig.UseAuth == true {
		var userPayload *UserPayload
		var cookie string
		var cErr error
		if decryptType == "cookie" {
			cookie, cErr = c.Cookie("session")
		} else {
			cookie = options["cookie"]
		}

		if cErr != nil {
			logger.Error(cErr)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error_code": http.StatusUnauthorized,
				"message": "Unauthorized, Access token is missing or invalid",
				"success": false,
			})
			c.Abort()
			return
		}

		signer := itsdangerous.NewTimestampSignature(
			appConfig.Secret,
			appConfig.Salt,
			appConfig.Sep,
			"hmac",
			nil,
			nil,
		)

		// unsign cookies to decrypt user info
		signed, err := signer.Unsign(cookie, 0)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Cannot unsign cookie",
				"error_code": http.StatusUnauthorized,
				"success": false,
			})
			c.Abort()
			return
		}

		// append == into string to convert to base64
		base64String := signed[1:len(signed)]
		base64String = base64String + strings.Repeat("=", (4 - (utf8.RuneCountInString(base64String) % 4)) % 4 )

		data, err := b64.URLEncoding.DecodeString(base64String)

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalidate base64 string",
				"error_code": http.StatusUnauthorized,
				"success": false,
			})
			c.Abort()
			return
		}

		b := bytes.NewReader(data)
		read, err := zlib.NewReader(b)

		defer read.Close()

		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error_code": http.StatusUnauthorized,
				"message": "Incorrect zlib data",
				"success": false,
			})
			c.Abort()
			return
		}

		payload, err := ioutil.ReadAll(read)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusExpectationFailed, gin.H{
				"message": "cannot read payload",
				"error_code": http.StatusExpectationFailed,
				"success": false,
			})
			c.Abort()
			return
		}
		logger.Debug(fmt.Sprint("payload decyprt is %s", string(payload)))

		if err := json.Unmarshal(payload, &userPayload); err != nil {
			logger.Error(err)
			c.JSON(http.StatusExpectationFailed, gin.H{
				"message": "cannot parse payload",
				"error_code": http.StatusExpectationFailed,
				"success": false,
			})
			c.Abort()
			return
		}
		// check expire of token
		expire := (*userPayload.Expires).Unix()

		if expire <= time.Now().Unix() {
			logger.Error(err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error_code": http.StatusUnauthorized,
				"message": "Token expire",
			})
			c.Abort()
			return
		}
		// set some value to use for project for context
		c.Set("globalConfig", *appConfig)
		c.Set("userId", userPayload.KsUserID)
		c.Set("tenantId", userPayload.TenantID)
		c.Set("tenantName", userPayload.TenantName)
		c.Set("email", userPayload.Email)
		c.Set("FullName", userPayload.FullName)
		c.Set("LastRegion", userPayload.LastRegion)
	}
	c.Next()
}