package model

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	FirstName  string
	MiddleName string
	LastName   string
	// Types       []string `json:"types,omitempty"`
	GradeLevel  string
	PhoneNumber string
	Email       string
}

// RoleIn ...
func isIn(val string, slice []string) bool {
	for _, item := range slice {
		if val == item {
			return true
		}
	}
	return false
}
