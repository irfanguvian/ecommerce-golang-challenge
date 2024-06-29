package controllers

import (
	"strconv"
	"sypnasis-golang-test-ecommerce/config"
	"sypnasis-golang-test-ecommerce/models"
	"sypnasis-golang-test-ecommerce/usecase"

	"github.com/gin-gonic/gin"
)

func CreateOrderHandler(c *gin.Context) {
	user, ok := c.Get("user")

	if !ok {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid user",
		})
		return
	}

	userObj, ok := user.(models.User) // Replace UserType with the actual type of the user object

	if !ok {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid user type",
		})
		return
	}

	OrderBody := usecase.OrderBody{
		UserID: userObj.ID,
	}

	_, err := usecase.CreateNewOrder(config.DB, OrderBody)

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Order created successfully",
	})
}

func SimulationPaymentHandler(c *gin.Context) {
	var body struct {
		ID uint
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	_, err := usecase.PaymentOrder(config.DB, usecase.OrderUpdateBody{
		ID: body.ID,
	})

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Order paid successfully",
	})
}

func GetOrderHandler(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid request params",
		})
		return
	}

	order, err := usecase.GetOrderByID(config.DB, uint(idInt))

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    order,
	})
}

func GetListOfUserOrderHandler(c *gin.Context) {
	limitQuery := c.DefaultQuery("limit", "10")
	offsetQuery := c.DefaultQuery("offset", "0")

	user, ok := c.Get("user")

	if !ok {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid user",
		})
		return
	}

	userObj, ok := user.(models.User) // Replace UserType with the actual type of the user object

	if !ok {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid user type",
		})
		return
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid request params",
		})
		return
	}

	offset, err := strconv.Atoi(offsetQuery)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid request params",
		})
		return
	}

	parameter := usecase.ParamaterOrderList{
		Limit:  limit,
		Offset: offset,
		UserID: userObj.ID,
	}

	orders, err := usecase.GetOrderList(config.DB, parameter)

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    orders,
	})
}
