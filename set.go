package godash

func Set(source interface{}, path string, value interface{}) interface{} {
	parts := ToPath(path)
	src := source

	// if source is not an array or map, no setting can take place
	if !IsArray(source) && !IsMap(source) {
		return source
	}

	for i, part := range parts {
		switch c := src.(type) {
		case []interface{}:
			// if the current path part is not an int, its invalid, return source
			if !IsInt(part) || part.(int) > len(c) - 1 {
				return source
			} else if i == len(parts) - 1 {
				if part.(int) < len(c) {
					c[part.(int)] = value
				}
				return source
			} else if len(c) > part.(int) {
				src = c[part.(int)]
				break
			}

			switch parts[i + 1].(type) {
			case int:
				c[part.(int)] = make([]interface{}, 0)
				src = c[part.(int)]
			case string:
				c[part.(int)] = make(map[string]interface{})
				src = c[part.(int)]
			default:
				return source
			}

		case map[string]interface{}:
			// if the current path part is not a string, its invalid, return source
			if !IsString(part) {
				return source
			} else if i == len(parts) - 1 {
				c[part.(string)] = value
				return source
			} else if value, ok := c[part.(string)]; ok {
				src = value
				break
			}

			// check the next path part for type and set it
			switch parts[i + 1].(type) {
			case int:
				c[part.(string)] = make([]interface{}, 0)
				src = c[part.(string)]
			case string:
				c[part.(string)] = make(map[string]interface{})
				src = c[part.(string)]
			default:
				return source
			}
		default:
			return source
		}
	}

	return source
}
