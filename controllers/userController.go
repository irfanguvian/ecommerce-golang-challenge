package controllers

import (
	"sypnasis-golang-test-ecommerce/config"
	"sypnasis-golang-test-ecommerce/usecase"

	"github.com/gin-gonic/gin"
)

func UserSignup(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	_, err := usecase.CreateNewUser(config.DB, body)

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "User created successfully",
	})
}

func UserLogin(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}

	token, err := usecase.LoginUser(config.RedisClient, config.DB, body)

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Login success",
		"token":   token,
	})
}

func UserAuthorize(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	user, err := usecase.AuthorizeJWT(config.RedisClient, config.DB, tokenString)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.Set("user", user)
	c.Next()
}
