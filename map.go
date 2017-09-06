package godash

import (
	"github.com/bhoriuchi/godash/internal"
)

// Map creates an array of values from the return value of an iteratee
func Map(collection interface{}, iteratee ...func(value interface{}, key interface{}) interface{}) []interface{} {
	m := make([]interface{}, 0)
	fn := internal.LoopIdentity

	if len(iteratee) > 0 {
		fn = iteratee[0]
	}

	ForEach(collection, func(value interface{}, key interface{}) {
		m = append(m, fn(value, key))
	})

	return m
}