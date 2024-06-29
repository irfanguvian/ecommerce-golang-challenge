package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	CategoryID  uint
	Price       uint64
	Stock       uint64
	Description string
	UserID      uint
}
