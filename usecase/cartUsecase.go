package usecase

import (
	"errors"
	"sypnasis-golang-test-ecommerce/models"
	"sypnasis-golang-test-ecommerce/repository"

	"gorm.io/gorm"
)

type CartBodyParameter struct {
	ProductID uint
	Quantity  uint64
	UserID    uint
}

func AddCart(DB *gorm.DB, cartBody CartBodyParameter) error {
	_, err := repository.GetProductByID(DB, cartBody.ProductID)

	if err != nil {
		return errors.New("product not found")
	}

	cart, err := repository.GetCartByUserIDAndProductID(DB, cartBody.UserID, cartBody.ProductID)

	if cart.ID == 0 && err != nil {
		cart := models.Cart{
			UserID:    cartBody.UserID,
			ProductID: cartBody.ProductID,
			Quantity:  cartBody.Quantity,
		}

		_, err := repository.CreateCart(DB, cart)

		if err != nil {
			return err
		}

		return nil
	}

	cart.Quantity += cartBody.Quantity

	repository.UpdateCart(DB, cart)

	return nil
}

func DecreaseCart(DB *gorm.DB, cartBody CartBodyParameter) error {
	_, err := repository.GetProductByID(DB, cartBody.ProductID)

	if err != nil {
		return errors.New("product not found")
	}

	cart, err := repository.GetCartByUserIDAndProductID(DB, cartBody.UserID, cartBody.ProductID)

	if cart.ID == 0 && err != nil {
		return errors.New("cart not found")
	}

	cart.Quantity -= cartBody.Quantity

	if cart.Quantity <= 0 {
		repository.DeleteCart(DB, cart.ID)
		return nil
	}

	repository.UpdateCart(DB, cart)

	return nil
}

func CartList(DB *gorm.DB, userID uint) ([]models.Cart, error) {
	carts, err := repository.GetCartByUserID(DB, userID)

	if err != nil {
		return nil, errors.New("cart not found")
	}

	return carts, nil
}
