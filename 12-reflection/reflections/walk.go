package reflections

import "reflect"

func walk(x any, fn func(string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Struct:
		for i := range val.NumField() {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice:
		for i := range val.Len() {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}

	// if val.Kind() == reflect.Slice {
	// 	for i := range val.Len() {
	// 		walk(val.Index(i).Interface(), fn)
	// 	}
	// 	return
	// }

	// for i := range val.NumField() {
	// 	field := val.Field(i)

	// 	switch field.Kind() {
	// 	case reflect.Struct:
	// 		for i := range val.NumField() {
	// 			walk(val.Field(i).Interface(), fn)
	// 		}
	// 	case reflect.Slice:
	// 		for i := range val.Len() {
	// 			walk(val.Index(i).Interface(), fn)
	// 		}
	// 	case reflect.String:
	// 		fn(field.String())
	// 	}

	// 	// val := reflect.ValueOf(x)

	// 	// if val.Kind() == reflect.Pointer {
	// 	// 	val = val.Elem()
	// 	// }

	// 	// if field.Kind() == reflect.String {
	// 	// 	fn(field.String())
	// 	// }

	// 	// if field.Kind() == reflect.Struct {
	// 	// 	walk(field.Interface(), fn)
	// 	// }

	// 	// if field.Kind() == reflect.String {
	// 	// 	fn(field.String())
	// 	// }
	// }

	// val := reflect.ValueOf(x)
	// field := val.Field(0)
	// fn(field.String())
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
