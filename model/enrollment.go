package model

import "github.com/jinzhu/gorm"

type Enrollment struct {
	gorm.Model

	PersonID uint `json:"personId"`
	Person   Person
	ClassID  uint `sql:"index"`
	TermID   uint `sql:"index"`
}
