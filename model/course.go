package model

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model
	Name        string   `json:"name"`
	GradeLevel  string   `json:"gradeLevel"`
	MaxStudents int      `json:"maxStudents"`
	Terms       []string `json:"terms,omitempty"`
}
