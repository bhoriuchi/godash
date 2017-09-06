package godash

import (
	"reflect"
)

func IsMap(value interface{}) bool {
	return reflect.TypeOf(value).Kind() == reflect.Map
}
