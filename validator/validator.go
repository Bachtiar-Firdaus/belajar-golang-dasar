package validator

import "reflect"

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			// return reflect.ValueOf(data).Field(i).Interface() != "" // cara singkat
			value := reflect.ValueOf(data).Field(i).Interface() // cara singkat
			if value == "" {
				return false
			}
		}
	}
	return true
}
