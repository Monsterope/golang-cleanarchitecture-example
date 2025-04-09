package utils

import (
	"reflect"
	"strings"
)

func CheckKeyIsHave(reqData interface{}) interface{} {
	valueOfReq := reflect.ValueOf(reqData).Elem()
	typeOfReq := reflect.TypeOf(reqData).Elem()

	datamap := map[string]interface{}{}

	for i := 0; i < valueOfReq.NumField(); i++ {
		v := valueOfReq.Field(i)
		k := typeOfReq.Field(i).Name
		t := v.Kind()
		switch t {
		case reflect.Int:
			if v.Int() != 0 {
				datamap[strings.ToLower(k)] = v.Int()
			}
		case reflect.String:
			if v.String() != "" {
				datamap[strings.ToLower(k)] = v.String()
			}
		}
	}

	return datamap
}
