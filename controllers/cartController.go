package controllers

import (
	"sypnasis-golang-test-ecommerce/config"
	"sypnasis-golang-test-ecommerce/models"
	"sypnasis-golang-test-ecommerce/usecase"

	"github.com/gin-gonic/gin"
)

func CartManageHandler(c *gin.Context) {
	var body struct {
		ManageType string `json:"manage_type"`
		Cart       struct {
			ProductID uint `json:"product_id"`
			Quantity  uint64
		}
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

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

	if body.Cart.ProductID == 0 || body.Cart.Quantity == 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid request body value send",
		})
		return
	}

	cartManageParam := usecase.CartBodyParameter{
		ProductID: body.Cart.ProductID,
		Quantity:  body.Cart.Quantity,
		UserID:    userObj.ID,
	}

	if body.ManageType == "ADD" {
		err := usecase.AddCart(config.DB, cartManageParam)

		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
	} else if body.ManageType == "DECREASE" {
		err := usecase.DecreaseCart(config.DB, cartManageParam)

		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
	} else {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Cart updated successfully",
	})
}

func CartListHandler(c *gin.Context) {
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

	carts, err := usecase.CartList(config.DB, userObj.ID)

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    carts,
	})
}
