package usecase

import (
	"errors"
	"sypnasis-golang-test-ecommerce/models"
	"sypnasis-golang-test-ecommerce/repository"

	"gorm.io/gorm"
)

type ProductBody struct {
	ID          uint
	Name        string
	CategoryID  uint
	Price       uint64
	Stock       uint64
	Description string
	UserID      uint
}

func CreateNewProduct(DB *gorm.DB, productBody ProductBody) (models.Product, error) {
	product := models.Product{
		Name:        productBody.Name,
		CategoryID:  productBody.CategoryID,
		Price:       productBody.Price,
		Stock:       productBody.Stock,
		Description: productBody.Description,
		UserID:      productBody.UserID,
	}
	return repository.CreateProduct(DB, product)
}

func UpdateProduct(DB *gorm.DB, productBody ProductBody) (*models.Product, error) {
	_, err := repository.GetProductByID(DB, productBody.ID)

	if err != nil {
		return nil, errors.New("product not found")
	}

	product := models.Product{
		Name:        productBody.Name,
		CategoryID:  productBody.CategoryID,
		Price:       productBody.Price,
		Stock:       productBody.Stock,
		Description: productBody.Description,
	}
	product.ID = productBody.ID

	return repository.UpdateProduct(DB, product)
}

func DeleteProduct(DB *gorm.DB, id uint) error {
	// checn is Product Exist
	_, err := repository.GetProductByID(DB, id)

	if err != nil {
		return errors.New("product not found")
	}

	return repository.DeleteProduct(DB, id)
}

type ParamaterProductList struct {
	Limit  int
	Offset int
	UserID uint
}

func GetProductList(DB *gorm.DB, parameter ParamaterProductList) ([]models.Product, error) {
	queryPrd := repository.QueryProduct{
		Limit:  parameter.Limit,
		Offset: parameter.Offset,
		UserID: parameter.UserID,
	}
	return repository.ListProducts(DB, queryPrd)
}

func GetProductByID(DB *gorm.DB, id uint) (models.Product, error) {
	return repository.GetProductByID(DB, id)
}
