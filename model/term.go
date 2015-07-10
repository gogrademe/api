package model

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mholt/binding"
)

type Term struct {
	gorm.Model
	Name       string     `json:"name"`
	SchoolYear SchoolYear `json:"schoolYear"`
	StartDate  time.Time  `json:"startDate"`
	EndDate    time.Time  `json:"endDate"`
}

func (t Term) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if t.Name == "" {
		errs = append(errs, RequiredErr("name"))
	}
	if t.StartDate.IsZero() {
		errs = append(errs, RequiredErr("startDate"))
	}
	if t.EndDate.IsZero() {
		errs = append(errs, RequiredErr("endDate"))
	}
	return errs
}

func (t *Term) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&t.ID:         "id",
		&t.Name:       "name",
		&t.SchoolYear: "schoolYear",
		&t.StartDate:  "startDate",
		&t.EndDate:    "endDate",
	}
}
