package repository

import (
	"sypnasis-golang-test-ecommerce/models"

	"gorm.io/gorm"
)

// CreateOrder adds a new order to the database.
func CreateOrderItemsBatch(DB *gorm.DB, orderItems []models.OrderItems) ([]models.OrderItems, error) {
	if err := DB.Create(&orderItems).Error; err != nil {
		return orderItems, err
	}
	return orderItems, nil
}
