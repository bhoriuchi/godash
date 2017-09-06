package godash

func Reduce(collection interface{}, iteratee func(accumulator interface{}, args ...interface{}) interface{}, accumulator ...interface{}) interface{} {
	var accum interface{}
	if len(accumulator) > 0 {
		 accum = accumulator[0]
	}

	switch c := collection.(type) {
	case []interface{}:
		if len(c) > 0 {
			accum = c[0]
		}
	case map[string]interface{}:
		keys := Keys(collection)
		if len(keys) > 0 {
			accum = c[keys[0].(string)]
		}
	case map[interface{}]interface{}:
		keys := Keys(collection)
		if len(keys) > 0 {
			accum = c[keys[0]]
		}
	default:
		return accum
	}

	ForEach(collection, func(value interface{}, key interface{}) {
		a := make([]interface{}, 0)
		a = append(a, value)
		a = append(a, key)
		a = append(a, collection)
		accum = iteratee(accum, a...)
	})

	return accum
}