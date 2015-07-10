package model

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/mholt/binding"
)

type Course struct {
	gorm.Model
	Name        string   `json:"name"`
	GradeLevel  string   `json:"gradeLevel"`
	MaxStudents int      `json:"maxStudents"`
	Terms       []string `json:"terms,omitempty"`
}

func (c Course) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if c.Name == "" {
		errs = append(errs, RequiredErr("name"))
	}
	if c.GradeLevel == "" {
		errs = append(errs, RequiredErr("gradeLevel"))
	}

	if len(c.Terms) <= 0 {
		errs = append(errs, RequiredErr("terms"))
	}
	return errs
}

func (c *Course) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&c.ID:         "id",
		&c.Name:       "name",
		&c.GradeLevel: "gradeLevel",
	}
}
