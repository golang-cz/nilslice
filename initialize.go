package nilslice

import (
	"reflect"
)

func Initialize(obj interface{}) interface{} {
	v := reflect.ValueOf(obj)
	initializeNils(v)

	return obj
}

func initializeNils(v reflect.Value) {
	// Dereference pointer(s).
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	if v.Kind() == reflect.Slice {
		// Initialize a nil slice.
		if v.IsNil() && v.CanSet() {
			v.Set(reflect.MakeSlice(v.Type(), 0, 0))
			return
		}

		// Recursively iterate over slice items.
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i)
			initializeNils(item)
		}
	}

	// Recursively iterate over struct fields.
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			initializeNils(field)
		}
	}
}
