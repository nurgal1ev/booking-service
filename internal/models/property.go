package models

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	Name          string
	Description   string
	Address       string
	City          string
	Country       string
	PricePerNight int
	PropertyType  string `gorm:"default:hotel"`

	Units []Unit `gorm:"foreignKey:PropertyID"`

	OwnerID uint
	Owner   User `gorm:"foreignKey:OwnerID"`
}
