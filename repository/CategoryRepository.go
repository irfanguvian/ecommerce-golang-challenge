package repository

import (
	"sypnasis-golang-test-ecommerce/models"

	"gorm.io/gorm"
)

func CreateCategory(DB *gorm.DB, Category models.Category) (models.Category, error) {
	if err := DB.Create(&Category).Error; err != nil {
		return Category, err
	}
	return Category, nil
}

func GetCategoryByID(DB *gorm.DB, id uint) (models.Category, error) {
	var Category models.Category
	if err := DB.First(&Category, id).Error; err != nil {
		return Category, err
	}
	return Category, nil
}

func UpdateCategory(DB *gorm.DB, Category models.Category) (*models.Category, error) {

	getCategory, err := GetCategoryByID(DB, Category.ID)

	if err != nil {
		return nil, err
	}

	if getCategory.ID == 0 {
		return nil, nil
	}

	result := DB.Model(getCategory).Updates(models.Category{Name: Category.Name})

	if result.Error != nil {
		return nil, result.Error
	}

	return &getCategory, nil
}

func DeleteCategory(DB *gorm.DB, id uint) error {
	var Category models.Category
	if err := DB.Delete(&Category, id).Error; err != nil {
		return err
	}
	return nil
}

func GetCategoryList(DB *gorm.DB) ([]models.Category, error) {
	var Category []models.Category
	if err := DB.Find(&Category).Error; err != nil {
		return Category, err
	}
	return Category, nil
}
