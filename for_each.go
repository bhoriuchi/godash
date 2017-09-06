package godash

import (
	"reflect"
)

// ForEach iterates over a collection executing a function on each iteration
func ForEach(collection interface{}, iteratee ...func(value interface{}, key interface{})) {
	fn := func(value interface{}, key interface{}) {
		return
	}

	if len(iteratee) > 0 {
		fn = iteratee[0]
	}

	switch reflect.TypeOf(collection).Kind() {
	case reflect.Map:
		m := reflect.ValueOf(collection)
		for _, key := range m.MapKeys() {
			fn(m.MapIndex(key), key)
		}
		break
	case reflect.Array, reflect.Slice:
		m := reflect.ValueOf(collection)
		for i := 0; i < m.Len(); i++ {
			fn(m.Index(i), i)
		}
		break
	default:
		break
	}
}
