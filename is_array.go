package godash

import (
	"reflect"
)

func IsArray(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.Array
}
