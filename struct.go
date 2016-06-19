package ng

import (
	"reflect"
	"strconv"
	"strings"
)

//IsStruct check an interface if its a struct or not
func IsStruct(i interface{}) (reflect.Value, bool) {
	v := reflect.ValueOf(i)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return v, false
	}

	return v, true
}

//ValidateStruct validating struct based on their field tags.
func ValidateStruct(i interface{}) *Errors {

	var (
		errs = NewErrors()
	)

	v, ok := IsStruct(i)
	if !ok {
		return errs
	}

	for index := 0; index < v.NumField(); index++ {
		val := v.Field(index)
		name := v.Type().Field(index).Name
		tags := v.Type().Field(index).Tag.Get("validate")
		var validations []string
		if tags != "" {
			validations = strings.Split(tags, ";")
		}

		if len(validations) == 0 {
			continue
		}

		for _, validation := range validations {
			var (
				acts    []string
				act     string
				errsVal *Errors
				fVal    string
				sVal    string
			)

			acts = strings.Split(validation, "=")
			act = acts[0]
			if len(acts) > 1 {
				for index, param := range strings.Split(acts[1], ",") {
					switch index {
					case 0:
						fVal = param
					case 1:
						sVal = param
					}
				}
			}

			switch act {
			case "required":
				errsVal = Required(name, val.Interface())
			case "min":
				param, _ := strconv.Atoi(fVal)
				errsVal = LengthMin(name, val.Interface().(string), param)
			case "max":
				param, _ := strconv.Atoi(fVal)
				errsVal = LengthMax(name, val.Interface().(string), param)
			case "range":
				param1, _ := strconv.Atoi(fVal)
				param2, _ := strconv.Atoi(sVal)
				errsVal = Range(name, val.Interface().(int), param1, param2)
			}

			errs.AppendErrors(errsVal)
		}
	}

	return errs
}
