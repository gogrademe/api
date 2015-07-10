package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type EmailConfirmation struct {
	gorm.Model
	UserID uint      `json:"userId"`
	UsedOn time.Time `json:"usedOn"`
}
