package controllers

import (
	"fmt"
	"strconv"
	"sypnasis-golang-test-ecommerce/config"
	"sypnasis-golang-test-ecommerce/models"
	"sypnasis-golang-test-ecommerce/usecase"

	"github.com/gin-gonic/gin"
)

func ProductManageHandler(c *gin.Context) {
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

	var body struct {
		ManageType string `json:"manage_type"`
		Product    struct {
			ID          uint
			Name        string
			CategoryID  uint `json:"category_id"`
			Price       uint64
			Stock       uint64
			Description string
		}
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	if body.Product.Name == "" || body.Product.CategoryID == 0 || body.Product.Price == 0 || body.Product.Stock == 0 || body.Product.Description == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid request body send",
		})
		return
	}

	if body.Product.CategoryID != 0 {
		_, err := usecase.GetCategory(config.DB, body.Product.CategoryID)

		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": "Category not found",
			})
			return
		}

	}

	productBody := usecase.ProductBody{
		ID:          body.Product.ID,
		Name:        body.Product.Name,
		CategoryID:  body.Product.CategoryID,
		Price:       body.Product.Price,
		Stock:       body.Product.Stock,
		Description: body.Product.Description,
		UserID:      userObj.ID,
	}

	if body.ManageType == "CREATE" {
		_, err := usecase.CreateNewProduct(config.DB, productBody)

		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
			"message": "Product created successfully",
		})
		return
	}

	if body.ManageType == "UPDATE" {
		if productBody.ID == 0 {
			c.JSON(400, gin.H{
				"success": false,
				"message": "product id required for update",
			})
			return
		}

		_, err := usecase.UpdateProduct(config.DB, productBody)

		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
			"message": "Product updated successfully",
		})
		return
	}

	if body.ManageType == "DELETE" {
		if productBody.ID == 0 {
			c.JSON(400, gin.H{
				"success": false,
				"message": "product id required for update",
			})
			return
		}

		err := usecase.DeleteProduct(config.DB, productBody.ID)

		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
			"message": "Product deleted successfully",
		})
		return
	}

	c.JSON(400, gin.H{
		"success": false,
		"message": "Invalid manage type",
	})
}

func GetProductListHandler(c *gin.Context) {

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

	parameter := usecase.ParamaterProductList{
		Limit:  limit,
		Offset: offset,
		UserID: userObj.ID,
	}

	products, err := usecase.GetProductList(config.DB, parameter)

	fmt.Println(products)

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success":  true,
		"products": products,
	})
}

func GetProductHandler(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid request params",
		})
		return
	}

	product, err := usecase.GetProductByID(config.DB, uint(idInt))

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"product": product,
	})
}
