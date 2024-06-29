package usecase

import (
	"errors"
	"sypnasis-golang-test-ecommerce/models"
	"sypnasis-golang-test-ecommerce/repository"

	"gorm.io/gorm"
)

type OrderBody struct {
	ID     uint
	UserID uint
}

type OrderUpdateBody struct {
	ID uint
}

func CreateNewOrder(DB *gorm.DB, OrderBody OrderBody) (models.Order, error) {
	var order models.Order
	// get cart
	carts, err := repository.GetCartByUserID(DB, OrderBody.UserID)

	if err != nil || len(carts) == 0 {
		return order, errors.New("cart not found")
	}
	// pupolate cart and get product id

	var productIDs []uint
	var cartIds []uint
	var QuantityProduct = make(map[uint]uint64)

	for _, val := range carts {
		productIDs = append(productIDs, val.ProductID)
		cartIds = append(cartIds, val.ID)
		QuantityProduct[val.ProductID] = val.Quantity
	}

	// get all product id
	queryProduct := repository.QueryProduct{
		ProductIds: productIDs,
		UserID:     OrderBody.UserID,
	}

	products, err := repository.ListProducts(DB, queryProduct)

	if err != nil || len(products) == 0 {
		return order, errors.New("product not found")
	}

	// process and count , bikin object sekaliatn buat order items
	var totalPrice uint64

	for i, val := range products {
		// quantity = quantity + OrderBody.Quantity
		totalPrice += (val.Price * QuantityProduct[val.ID])
		if val.Stock < QuantityProduct[val.ID] {
			return order, errors.New("quantity not enough")
		}
		products[i].Stock = val.Stock - QuantityProduct[val.ID]
		_, err := repository.UpdateProduct(DB, products[i])
		if err != nil {
			return order, errors.New("error update product")

		}
	}
	// bikin order
	order.UserID = OrderBody.UserID
	order.TotalPrice = totalPrice
	order.Status = 0

	order, err = repository.CreateOrder(DB, order)

	if err != nil {
		return order, errors.New("error create order items")
	}

	var orderItems []models.OrderItems

	for _, val := range products {
		orderItem := models.OrderItems{
			OrderID:   order.ID,
			ProductID: val.ID,
			Name:      val.Name,
			Quantity:  QuantityProduct[val.ID],
			Price:     val.Price * QuantityProduct[val.ID],
		}
		orderItems = append(orderItems, orderItem)
	}

	// bikin order items

	_, err = repository.CreateOrderItemsBatch(DB, orderItems)

	if err != nil {
		return order, errors.New("error create order items")
	}

	// delete cart

	err = repository.DeleteCartBatch(DB, cartIds)

	if err != nil {
		return order, errors.New("error create order items")
	}

	return order, nil
}

func PaymentOrder(DB *gorm.DB, OrderUpdateBody OrderUpdateBody) (*models.Order, error) {
	order, err := repository.GetOrderById(DB, OrderUpdateBody.ID)

	if err != nil {
		return nil, errors.New("order not found")
	}

	if order.Status == 1 {
		return nil, errors.New("order already paid")
	}

	order.Status = 1

	return repository.UpdateOrder(DB, order)
}

type ParamaterOrderList struct {
	Limit  int
	Offset int
	UserID uint
}

func GetOrderList(DB *gorm.DB, parameter ParamaterOrderList) ([]models.Order, error) {
	queryOrd := repository.QueryOrder{
		Limit:  parameter.Limit,
		Offset: parameter.Offset,
		UserID: parameter.UserID,
	}
	return repository.ListOrders(DB, queryOrd)
}

func GetOrderByID(DB *gorm.DB, id uint) (models.Order, error) {
	return repository.GetOrderById(DB, id)
}
