package godash

// Keys returns an array of keys from a collection
func Keys(collection interface{}) []interface{} {
	k := make([]interface{}, 0)

	ForEach(collection, func(value interface{}, key interface{}) {
		k = append(k, key)
	})

	return k
}