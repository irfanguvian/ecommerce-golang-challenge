package repository

import (
	"sypnasis-golang-test-ecommerce/models"

	"gorm.io/gorm"
)

// CreateOrder adds a new order to the database.
func CreateOrder(DB *gorm.DB, order models.Order) (models.Order, error) {
	if err := DB.Create(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

// GetOrderById retrieves a single order by its ID.
func GetOrderById(DB *gorm.DB, id uint) (models.Order, error) {
	var order models.Order
	if err := DB.First(&order, id).Error; err != nil {
		return order, err
	}
	return order, nil
}

// UpdateOrder updates an existing order's details.
func UpdateOrder(DB *gorm.DB, order models.Order) (*models.Order, error) {
	if err := DB.Save(&order).Error; err != nil {
		return &order, err
	}
	return &order, nil
}

// DeleteOrder removes an order from the database by its ID.
func DeleteOrder(DB *gorm.DB, id uint) error {
	if err := DB.Delete(&models.Order{}, id).Error; err != nil {
		return err
	}
	return nil
}

type QueryOrder struct {
	Offset int
	Limit  int
	UserID uint
}

// ListOrders retrieves all orders from the database.
func ListOrders(DB *gorm.DB, QueryOrder QueryOrder) ([]models.Order, error) {
	var orders []models.Order

	if QueryOrder.UserID != 0 {
		DB = DB.Where("user_id = ?", QueryOrder.UserID)
	}

	if QueryOrder.Offset != 0 {
		DB = DB.Offset(QueryOrder.Limit)
	}

	if QueryOrder.Limit != 0 {
		DB = DB.Limit(QueryOrder.Limit)
	}
	if err := DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
