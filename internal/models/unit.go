package models

import "gorm.io/gorm"

type Unit struct {
	gorm.Model
	Name          string
	Description   string
	PricePerNight int
	Capacity      int
	IsAvailable   bool `gorm:"default:true"`

	PropertyID uint
	Property   Property `gorm:"foreignKey:PropertyID"`
}
