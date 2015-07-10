package model

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/mholt/binding"
)

type Person struct {
	gorm.Model
	FirstName  string `json:"firstName,omitempty"`
	MiddleName string `json:"middleName,omitempty"`
	LastName   string `json:"lastName,omitempty"`
	// Types       []string `json:"types,omitempty"`
	GradeLevel  string `json:"gradeLevel"`
	PhoneNumber string `json:"personId"`
	Email       string `json:"email"`
}

// type PersonRes struct {
// 	Person
// 	Account User
// }

// RoleIn ...
func isIn(val string, slice []string) bool {
	for _, item := range slice {
		if val == item {
			return true
		}
	}
	return false
}

func (p Person) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if p.FirstName == "" {
		errs = append(errs, RequiredErr("firstName"))
	}
	if p.LastName == "" {
		errs = append(errs, RequiredErr("lastName"))
	}
	// if len(p.Types) == 0 {
	// 	errs = append(errs, RequiredErr("types"))
	// }
	// if isIn("Student", p.Types) && p.GradeLevel == "" {
	// 	errs = append(errs, RequiredErr("gradeLevel"))
	// }

	return errs
}

func (p *Person) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&p.ID:         "id",
		&p.FirstName:  "firstName",
		&p.MiddleName: "middleName",
		&p.LastName:   "lastName",
		// &p.Types:       "types",
		&p.GradeLevel:  "gradeLevel",
		&p.PhoneNumber: "phoneNumber",
		&p.Email:       "email",
	}
}
