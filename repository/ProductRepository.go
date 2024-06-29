package repository

import (
	"sypnasis-golang-test-ecommerce/models"

	"gorm.io/gorm"
)

func GetProductByID(DB *gorm.DB, id uint) (models.Product, error) {
	var product models.Product
	if err := DB.First(&product, id).Error; err != nil {
		return product, err
	}
	return product, nil
}

func CreateProduct(DB *gorm.DB, product models.Product) (models.Product, error) {
	if err := DB.Create(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func UpdateProduct(DB *gorm.DB, product models.Product) (*models.Product, error) {
	if err := DB.Save(&product).Error; err != nil {
		return &product, err
	}
	return &product, nil
}

func DeleteProduct(DB *gorm.DB, id uint) error {
	if err := DB.Delete(&models.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}

type QueryProduct struct {
	Offset     int
	Limit      int
	UserID     uint
	ProductIds []uint
}

func ListProducts(DB *gorm.DB, queryProduct QueryProduct) ([]models.Product, error) {
	var products []models.Product

	if queryProduct.Offset != 0 {
		DB = DB.Offset(queryProduct.Limit)
	}

	if queryProduct.Limit != 0 {
		DB = DB.Limit(queryProduct.Limit)
	}

	if queryProduct.UserID != 0 {
		DB = DB.Where("user_id != ?", queryProduct.UserID)
	}

	if len(queryProduct.ProductIds) > 0 {
		DB = DB.Where("id in ?", queryProduct.ProductIds)
	}

	err := DB.
		Find(&products).Error

	if err != nil {
		return nil, err
	}
	return products, nil
}
