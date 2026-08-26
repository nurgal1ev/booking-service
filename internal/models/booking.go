package models

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	UserID uint
	User   User `gorm:"foreignkey:UserID"`

	PropertyID uint
	Property   Property `gorm:"foreignkey:PropertyID"`

	Price int
}
