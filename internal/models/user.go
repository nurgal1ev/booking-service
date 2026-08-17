package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string
	Email     string `gorm:"unique"`
	Password  string
	Role      string
	UpdatedAt *time.Time
}
