package validator

import (
	"errors"
	"github.com/gogf/gf/util/gvalid"
	"log"
	"math"
	"reflect"
)

func InitRequiredValidator() {
	err := gvalid.RegisterRule("required", func(rule string, value interface{}, message string, params map[string]interface{}) error {
		reflectValue := reflect.ValueOf(value)
		if reflectValue.Kind() == reflect.Ptr {
			reflectValue = reflectValue.Elem()
		}
		isEmpty := false
		switch reflectValue.Kind() {
		case reflect.Bool:
			isEmpty = !reflectValue.Bool()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			isEmpty = reflectValue.Int() == 0
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			isEmpty = reflectValue.Uint() == 0
		case reflect.Float32, reflect.Float64:
			isEmpty = math.Float64bits(reflectValue.Float()) == 0
		case reflect.Complex64, reflect.Complex128:
			c := reflectValue.Complex()
			isEmpty = math.Float64bits(real(c)) == 0 && math.Float64bits(imag(c)) == 0
		case reflect.String, reflect.Map, reflect.Array, reflect.Slice:
			isEmpty = reflectValue.Len() == 0
		}

		if isEmpty {
			return errors.New(message)
		}

		return nil
	})

	if err != nil {
		log.Fatal("Gvalid Register Rule Error:" + err.Error())
	}
}
