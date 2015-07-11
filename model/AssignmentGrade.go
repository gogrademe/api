package model

import "github.com/jinzhu/gorm"

type Attempt struct {
	gorm.Model
	AssignmentID uint    `sql:"index"`
	PersonID     uint    `sql:"index"`
	Score        string  `json:"score"`
	GradeAverage float32 `json:"gradeAverage"`
}

// AttemptResource is used for in the API to add new attempts.
// type AttemptResource struct {
// 	AssignmentID uint `json:"assignmentId"`
// 	PersonID     uint `json:"personId"`
// 	Attempt
// }

type AssignmentAttempts struct {
	gorm.Model
	AssignmentID   uint       `json:"assignmentId"`
	PersonID       uint       `json:"personId"`
	LatestAttempt  Attempt    `json:"latestAttempt"`
	AttemptHistory []*Attempt `json:"attemptHistory"`
}

type GradebookResource struct {
	Enrollment
	AssignmentAttempts []AssignmentAttempts `json:"assignmentAttempts"`
}

// FieldMap ...
// func (a *AttemptResource) FieldMap() binding.FieldMap {
// 	return binding.FieldMap{
//
// 		&a.AssignmentID: "assignmentId",
// 		&a.PersonID:     "personId",
// 		&a.Score:        "score",
// 	}
// }
