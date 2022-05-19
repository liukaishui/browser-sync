package utils

func InSlice(value string, vSlice []string) bool {
	for _, v := range vSlice {
		if v == value {
			return true
		}
	}
	return false
}
