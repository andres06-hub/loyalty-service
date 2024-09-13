package parsers

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

type ParamsParser struct{}

func Params() Parser {
	return &ParamsParser{}
}

func (*ParamsParser) Parse(ctx *http.Request, schema any) (err error) {
	if err = ctx.ParseForm(); err != nil {
		return err
	}

	v := reflect.ValueOf(schema).Elem()
	numFields := v.NumField()

	for i := 0; i < numFields; i++ {
		field := v.Field(i)
		fieldName := v.Type().Field(i).Name
		formValue := ctx.FormValue(fieldName)

		if formValue == "" {
			continue
		}

		if field.CanSet() {
			switch field.Kind() {
			case reflect.String:
				field.SetString(formValue)
			case reflect.Int:
				intValue, err := strconv.Atoi(formValue)
				if err != nil {
					return fmt.Errorf("invalid value for field %s: %v", fieldName, err)
				}
				field.SetInt(int64(intValue))
			default:
				return fmt.Errorf("unsupported field type: %s", field.Kind())
			}
		} else {
			return fmt.Errorf("can't set field value: %s", fieldName)
		}
	}

	return nil
}
