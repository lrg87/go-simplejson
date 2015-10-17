package simplejson

import (
	"encoding/json"
	"errors"
	"reflect"
)

func generateStringKey(key interface{}) (string, error) {
	if key == nil {
		return "", errNil
	}
	switch reflect.ValueOf(key).Type().Kind() {
	case reflect.Bool:
		fallthrough
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
		fallthrough
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		str, err := json.Marshal(key)
		if err == nil {
			return string(str), nil
		}
	case reflect.String:
		return key.(string), nil
	}
	return "", errors.New("Type is error")
}
