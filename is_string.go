package godash

func IsString(value interface{}) bool {
	switch value.(type) {
	case string:
		return true
	default:
		return false
	}
}
