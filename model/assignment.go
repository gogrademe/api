package model

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mholt/binding"
)

//Assignment ...
type Assignment struct {
	gorm.Model
	Name     string          `json:"name"`
	ClassID  uint            `json:"classId"`
	TermID   uint            `json:"termId"`
	GroupID  uint            `json:"groupId"`
	Group    AssignmentGroup `json:"group"`
	MaxScore int16           `json:"maxScore"`
	DueDate  time.Time       `json:"dueDate"`
}

// FieldMap ...
func (a *Assignment) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&a.ID:       "id",
		&a.Name:     "name",
		&a.GroupID:  "groupId",
		&a.ClassID:  "classId",
		&a.MaxScore: "maxScore",
		&a.TermID:   "termId",
		&a.DueDate:  "dueDate",
	}
}

func (a Assignment) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if a.Name == "" {
		errs = append(errs, RequiredErr("name"))
	}
	if a.MaxScore <= 0 {
		errs = append(errs, RequiredErr("maxScore"))
	}
	return errs
}
