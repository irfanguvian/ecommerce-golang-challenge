package main

import (
	"sypnasis-golang-test-ecommerce/config"
	"sypnasis-golang-test-ecommerce/models"
)

func init() {
	config.LoadConfig()
	config.ConnectionPostgres(config.ConfigENV)
	config.NewRedisClient(config.ConfigENV)
}

func main() {
	config.DB.AutoMigrate(&models.AccessToken{})
	config.DB.AutoMigrate(&models.Cart{})
	config.DB.AutoMigrate(&models.Category{})
	config.DB.AutoMigrate(&models.Order{})
	config.DB.AutoMigrate(&models.OrderItems{})
	config.DB.AutoMigrate(&models.Product{})
	config.DB.AutoMigrate(&models.User{})
}
