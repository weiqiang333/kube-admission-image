package method

// FindsStringSlice: 查找 String Slice 值是否存在
func FindsStringSlice(slice []string, value string) bool {
	if len(slice) == 0 {
		return false
	}
	if slice[0] == value {
		return true
	}
	return FindsStringSlice(slice[1:], value)
}
