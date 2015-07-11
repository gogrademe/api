package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Announcement ...
type Announcement struct {
	gorm.Model
	Name       string
	PersonID   uint `sql:"index"`
	Author     Person
	PostedDate time.Time
}
