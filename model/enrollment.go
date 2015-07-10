package model

import (
	"github.com/jinzhu/gorm"
	"github.com/mholt/binding"
)

type Enrollment struct {
	gorm.Model

	PersonID uint `json:"personId"`
	Person   Person
	ClassID  uint `json:"classId"`
	TermID   uint `json:"termId"`
}

func (e *Enrollment) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&e.ID:       "id",
		&e.PersonID: "personId",
		&e.TermID:   "termId",
		&e.ClassID:  "classId",
	}
}
