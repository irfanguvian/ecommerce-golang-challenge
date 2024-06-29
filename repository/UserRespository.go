package repository

import (
	"sypnasis-golang-test-ecommerce/models"

	"gorm.io/gorm"
)

func GetUserByID(DB *gorm.DB, id int) (models.User, error) {
	var user models.User
	if err := DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(DB *gorm.DB, user models.User) (models.User, error) {
	if err := DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByEmail(DB *gorm.DB, email string) (models.User, error) {
	var user models.User
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
