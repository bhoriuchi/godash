package godash

// Get gets the value at a deeply nested path
func Get(source interface{}, path string, defaultValue ...interface{}) interface{} {
	var defVal interface{}
	if len(defaultValue) > 0 {
		defVal = defaultValue[0]
	} else {
		defVal = nil
	}
	parts := ToPath(path)
	for i, part := range parts {
		switch c := source.(type) {
		case []interface{}:
			switch part.(type) {
			case int:
				if part.(int) < len(c) {
					source = c[part.(int)]
					if i == len(parts) - 1 {
						return source
					}
				} else {
					return defVal
				}
			default:
				return defVal
			}
			break
		case map[string]interface{}:
			switch part.(type) {
			case string:
				if value, ok := c[part.(string)]; ok {
					source = value
					if i == len(parts) - 1 {
						return source
					}
				} else {
					return defVal
				}
			default:
				return defVal
			}

			break
		default:
			break
		}
	}

	return defVal
}