package usecase

import (
	"errors"
	"sypnasis-golang-test-ecommerce/models"
	"sypnasis-golang-test-ecommerce/repository"

	"gorm.io/gorm"
)

type CategoryBody struct {
	ID   uint
	Name string
}

func CreateNewCategory(DB *gorm.DB, categoryBody CategoryBody) (models.Category, error) {
	category := models.Category{
		Name: categoryBody.Name,
	}
	return repository.CreateCategory(DB, category)
}

func UpdateCategory(DB *gorm.DB, categoryBody CategoryBody) (*models.Category, error) {
	_, err := repository.GetCategoryByID(DB, categoryBody.ID)

	if err != nil {
		return nil, errors.New("product not found")
	}

	category := models.Category{
		Name: categoryBody.Name,
	}

	category.ID = categoryBody.ID

	return repository.UpdateCategory(DB, category)
}

func DeleteCategory(DB *gorm.DB, id uint) error {
	_, err := repository.GetCategoryByID(DB, id)

	if err != nil {
		return errors.New("category not found")
	}

	return repository.DeleteCategory(DB, id)
}

func GetCategoryList(DB *gorm.DB) ([]models.Category, error) {
	return repository.GetCategoryList(DB)
}

func GetCategory(DB *gorm.DB, id uint) (models.Category, error) {
	return repository.GetCategoryByID(DB, id)
}
