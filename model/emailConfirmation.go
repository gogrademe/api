package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type EmailConfirmation struct {
	gorm.Model
	UserID uint      `sql:"index"`
	UsedOn time.Time `json:"usedOn"`
}
