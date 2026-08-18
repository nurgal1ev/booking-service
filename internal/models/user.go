package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Username  string
	Email     string `gorm:"unique"`
	Password  string `gorm:"unique"`
	Role      string
	UpdatedAt *time.Time
}
