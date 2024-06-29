package repository

import (
	"sypnasis-golang-test-ecommerce/models"

	"gorm.io/gorm"
)

// CreateCart adds a new cart to the database.
func CreateCart(DB *gorm.DB, cart models.Cart) (models.Cart, error) {
	if err := DB.Create(&cart).Error; err != nil {
		return models.Cart{}, err
	}
	return cart, nil
}

// GetCartByID retrieves a single cart by its ID.
func GetCartByID(DB *gorm.DB, id uint) (models.Cart, error) {
	var cart models.Cart
	if err := DB.First(&cart, id).Error; err != nil {
		return models.Cart{}, err
	}
	return cart, nil
}

// UpdateCart updates an existing cart's details.
func UpdateCart(DB *gorm.DB, cart models.Cart) (*models.Cart, error) {
	if err := DB.Save(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

// DeleteCart removes a cart from the database by its ID.
func DeleteCart(DB *gorm.DB, id uint) error {
	if err := DB.Delete(&models.Cart{}, id).Error; err != nil {
		return err
	}
	return nil
}

// ListCarts retrieves all carts from the database.
func ListCarts(DB *gorm.DB) ([]models.Cart, error) {
	var carts []models.Cart
	if err := DB.Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func GetCartByUserID(DB *gorm.DB, userID uint) ([]models.Cart, error) {
	var carts []models.Cart
	err := DB.Joins("Product").Where("carts.user_id = ?", userID).Find(&carts)
	return carts, err.Error
}

func GetCartByUserIDAndProductID(DB *gorm.DB, userID uint, productID uint) (models.Cart, error) {
	var cart models.Cart
	if err := DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error; err != nil {
		return models.Cart{}, err
	}
	return cart, nil
}

func DeleteCartBatch(DB *gorm.DB, id []uint) error {
	if err := DB.Where("id in ?", id).Delete(&models.Cart{}).Error; err != nil {
		return err
	}
	return nil
}
