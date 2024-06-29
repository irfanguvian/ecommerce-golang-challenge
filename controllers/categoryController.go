package controllers

import (
	"sypnasis-golang-test-ecommerce/config"
	"sypnasis-golang-test-ecommerce/usecase"

	"github.com/gin-gonic/gin"
)

func CategoryManageHandler(c *gin.Context) {
	var body struct {
		ManageType string `json:"manage_type"`
		category   struct {
			ID   uint
			Name string
		}
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	if body.ManageType == "CREATE" {
		_, err := usecase.CreateNewCategory(config.DB, body.category)

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
		_, err := usecase.UpdateCategory(config.DB, body.category)

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
		err := usecase.DeleteCategory(config.DB, body.category.ID)

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
