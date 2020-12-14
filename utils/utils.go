package utils

// Contain checks value was included slice
func Contain(s interface{}, v interface{}) bool {
	switch slice := s.(type) {
	case []string:
		for _, val := range slice {
			if val == v {
				return true
			}
		}
	case []int:
		for _, val := range slice {
			if val == v {
				return true
			}
		}
	}
	return false
}
