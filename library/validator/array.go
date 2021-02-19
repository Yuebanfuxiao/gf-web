package validator

import (
	"errors"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
	"log"
	"reflect"
	"strconv"
	"strings"
)

var (
	typeMap = map[string]struct{}{
		"integer": {},
		"string":  {},
		"float":   {},
		"boolean": {},
		"any":     {},
	}
)

// array check
// array:integer,min-count,max-count,min-val,max-val
// array:string,min-count,max-count,min-len,max-len
// array:float,min-count,max-count,min-val,max-val
// array:boolean,min-count,max-count
// array:any,min-count,max-count
func InitArrayValidator() {
	err := gvalid.RegisterRule("array", func(rule string, value interface{}, message string, params map[string]interface{}) error {
		reflectValue := reflect.ValueOf(value)

		if reflectValue.Kind() == reflect.Ptr {
			reflectValue = reflectValue.Elem()
		}

		if kind := reflectValue.Kind(); kind == reflect.Array || kind == reflect.Slice {
			if rulePattern := strings.Split(rule, ":")[1]; rulePattern != "" {
				patterns := strings.Split(rulePattern, ",")

				if patternsLen := len(patterns); patternsLen >= 1 {
					if patternsLen >= 2 {
						if minCount := gconv.Int(patterns[1]); minCount >= 0 {
							if count := len(value.([]interface{})); count < minCount {
								return errors.New(message)
							}
						}
					}

					if patternsLen >= 3 {
						if maxCount := gconv.Int(patterns[2]); maxCount >= 0 {
							if count := len(value.([]interface{})); count > maxCount {
								return errors.New(message)
							}
						}
					}

					if _, ok := typeMap[patterns[0]]; ok {
						if patterns[0] != "any" {
							for _, v := range value.([]interface{}) {
								switch patterns[0] {
								case "integer":
									if n, err := strconv.Atoi(gconv.String(v)); err != nil {
										return errors.New(message)
									} else {
										if patternsLen >= 3 && patterns[2] != "" && n < gconv.Int(patterns[2]) {
											return errors.New(message)
										}

										if patternsLen >= 4 && patterns[3] != "" && n < gconv.Int(patterns[3]) {
											return errors.New(message)
										}
									}

								case "string":
									n := len(gconv.String(v))

									if patternsLen >= 3 && patterns[2] != "" && n < gconv.Int(patterns[2]) {
										return errors.New(message)
									}

									if patternsLen >= 4 && patterns[3] != "" && n < gconv.Int(patterns[3]) {
										return errors.New(message)
									}

								case "float":
									if n, err := strconv.ParseFloat(gconv.String(v), 10); err != nil {
										return errors.New(message)
									} else {
										if patternsLen >= 3 && patterns[2] != "" && n < gconv.Float64(patterns[2]) {
											return errors.New(message)
										}

										if patternsLen >= 4 && patterns[3] != "" && n < gconv.Float64(patterns[3]) {
											return errors.New(message)
										}
									}

								case "boolean":
									if gconv.String(v) != "true" && gconv.String(v) != "false" {
										return errors.New(message)
									}
								}
							}
						}
					}
				}
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal("Gvalid Register Rule Error:" + err.Error())
	}
}
