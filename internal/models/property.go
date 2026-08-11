package models

import "gorm.io/gorm"

type Properties struct {
	gorm.Model
	Title       string
	Description string
	Address     string
	Price       int
}
