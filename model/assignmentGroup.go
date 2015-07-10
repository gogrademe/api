package model

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/mholt/binding"
)

// AssignmentGroup ...
type AssignmentGroup struct {
	gorm.Model
	Name    string  `json:"name"`
	Weight  float64 `json:"weight"`
	ClassID uint    `json:"classId"`
	TermID  uint    `json:"termId"`
}

// FieldMap ...
func (a *AssignmentGroup) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&a.ID:      "id",
		&a.Name:    "name",
		&a.Weight:  "weight",
		&a.ClassID: "classId",
		&a.TermID:  "termId",
	}
}

// Validate ...
func (a AssignmentGroup) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if a.Name == "" {
		errs = append(errs, RequiredErr("name"))
	}

	if a.Weight > 1 || a.Weight < 0.005 {
		errs = append(errs, binding.Error{
			FieldNames: []string{"weight"},
			Message:    "must be between .5% and 100%",
		})
	}
	return errs
}
