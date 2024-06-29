package main

import (
	"sypnasis-golang-test-ecommerce/config"
	"sypnasis-golang-test-ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfig()
	config.ConnectionPostgres(config.ConfigENV)
	config.NewRedisClient(config.ConfigENV)
}

func main() {
	// Initialize Gin
	router := gin.Default()

	// User routes
	router.POST("/signup", controllers.UserSignup)
	router.POST("/login", controllers.UserLogin)

	// Product routes
	router.GET("/products", controllers.UserAuthorize, controllers.GetProductListHandler)
	router.GET("/product/:id", controllers.UserAuthorize, controllers.GetProductHandler)
	router.POST("/product", controllers.UserAuthorize, controllers.ProductManageHandler)

	// Category routes
	router.POST("/category", controllers.UserAuthorize, controllers.CategoryManageHandler)

	// Order routes
	router.POST("/order", controllers.UserAuthorize, controllers.CreateOrderHandler)
	router.GET("/orders", controllers.UserAuthorize, controllers.GetListOfUserOrderHandler)
	router.GET("/order/:id", controllers.UserAuthorize, controllers.GetOrderHandler)

	// // Payment route
	router.POST("/payments", controllers.UserAuthorize, controllers.SimulationPaymentHandler)

	// // Cart routes
	router.POST("/cart", controllers.UserAuthorize, controllers.CartManageHandler)
	router.GET("/carts", controllers.UserAuthorize, controllers.CartListHandler)

	// Start the server on port 3000
	port := ":" + config.ConfigENV.PORT
	router.Run(port)
}
