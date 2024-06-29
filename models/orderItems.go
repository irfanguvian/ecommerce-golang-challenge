package models

import "gorm.io/gorm"

type OrderItems struct {
	gorm.Model
	OrderID   uint
	Name      string
	Quantity  uint64
	ProductID uint
	Price     uint64
}
