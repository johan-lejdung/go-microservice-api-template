package middleware

// StringInSlice will return true if string is in slice
func StringInSlice(s string, slice []string) bool {
	for _, elem := range slice {
		if s == elem {
			return true
		}
	}

	return false
}
