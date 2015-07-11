package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Assignment ...
type Assignment struct {
	gorm.Model
	Name     string          `json:"name"`
	ClassID  uint            `sql:"index"`
	TermID   uint            `sql:"index"`
	GroupID  uint            `sql:"index"`
	Group    AssignmentGroup `json:"group"`
	MaxScore int16           `json:"maxScore"`
	DueDate  time.Time       `json:"dueDate"`
}
