package model

import "github.com/jinzhu/gorm"

// AssignmentGroup ...
type AssignmentGroup struct {
	gorm.Model
	Name    string  `json:"name"`
	Weight  float64 `json:"weight"`
	ClassID uint    `sql:"index"`
	TermID  uint    `sql:"index"`
}
