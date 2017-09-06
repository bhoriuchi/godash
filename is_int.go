package godash

func IsInt(value interface{}) bool {
	switch value.(type) {
	case int:
		return true
	default:
		return false
	}
}
