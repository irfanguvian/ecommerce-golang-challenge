package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID     uint
	TotalPrice uint64
	Status     uint8 // 0: pending, 1: paid
}
