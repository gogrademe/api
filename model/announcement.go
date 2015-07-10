package model

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mholt/binding"
)

//Announcement ...
type Announcement struct {
	gorm.Model
	Name       string
	PersonID   uint
	Author     Person
	PostedDate time.Time
}

// FieldMap ...
func (a *Announcement) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&a.ID:   "id",
		&a.Name: "name",
	}
}
func (a Announcement) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if a.Name == "" {
		errs = append(errs, RequiredErr("name"))
	}
	if a.PersonID == 0 {
		errs = append(errs, RequiredErr("personId"))
	}
	return errs
}
