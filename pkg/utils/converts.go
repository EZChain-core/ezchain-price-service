package utils

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
	"strings"
	"crypto/rand"
)

// generate random token with 64 character
func TokenGenerator() string {
	b := make([]byte, 64)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// convert from struct to map
func StructToMap(st interface{}) map[string]interface{} {

	reqRules := make(map[string]interface{})

	v := reflect.ValueOf(st)
	t := reflect.TypeOf(st)

	for i := 0; i < v.NumField(); i++ {
		key := strings.ToLower(t.Field(i).Name)
		typ := v.FieldByName(t.Field(i).Name).Kind().String()
		structTag := t.Field(i).Tag.Get("json")
		jsonName := strings.TrimSpace(strings.Split(structTag, ",")[0])
		value := v.FieldByName(t.Field(i).Name)

		// if jsonName is not empty use it for the key
		if jsonName != "" && jsonName != "-" {
			key = jsonName
		}

		if typ == "string" {
			if !(value.String() == "" && strings.Contains(structTag, "omitempty")) {
				fmt.Println(key, value)
				fmt.Println(key, value.String())
				reqRules[key] = value.String()
			}
		} else if typ == "int" {
			reqRules[key] = value.Int()
		} else {
			reqRules[key] = value.Interface()
		}

	}

	return reqRules
}

//func StructToMap(i interface{}) (values url.Values) {
//	values = url.Values{}
//	iVal := reflect.ValueOf(i).Elem()
//	typ := iVal.Type()
//	for i := 0; i < iVal.NumField(); i++ {
//		f := iVal.Field(i)
//		// You ca use tags here...
//		// tag := typ.Field(i).Tag.Get("tagname")
//		// Convert each type into a string for the url.Values string map
//		var v string
//		switch f.Interface().(type) {
//		case int, int8, int16, int32, int64:
//			v = strconv.FormatInt(f.Int(), 10)
//		case uint, uint8, uint16, uint32, uint64:
//			v = strconv.FormatUint(f.Uint(), 10)
//		case float32:
//			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
//		case float64:
//			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
//		case []byte:
//			v = string(f.Bytes())
//		case string:
//			v = f.String()
//		}
//		values.Set(typ.Field(i).Name, v)
//	}
//	return
//}