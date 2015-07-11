package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Term struct {
	gorm.Model
	Name       string     `json:"name"`
	SchoolYear SchoolYear `json:"schoolYear"`
	StartDate  time.Time  `json:"startDate"`
	EndDate    time.Time  `json:"endDate"`
}
