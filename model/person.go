package model

// RoleIn ...
func isIn(val string, slice []string) bool {
	for _, item := range slice {
		if val == item {
			return true
		}
	}
	return false
}
