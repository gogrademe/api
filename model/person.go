package model

type Role int

const (
	IsAdmin Role = 1 << iota
	IsTeacher
	IsParent
	IsStudent
)

// RoleIn ...
func isIn(val string, slice []string) bool {
	for _, item := range slice {
		if val == item {
			return true
		}
	}
	return false
}
