package repository

import (
	"sypnasis-golang-test-ecommerce/models"

	"gorm.io/gorm"
)

func GetUserByAccessId(DB *gorm.DB, id string) (models.AccessToken, error) {
	var AccessToken models.AccessToken
	if err := DB.First(&AccessToken, "access_id = ?", id).Error; err != nil {
		return AccessToken, err
	}
	return AccessToken, nil
}

func CreateAccessToken(DB *gorm.DB, accessToken models.AccessToken) (models.AccessToken, error) {
	if err := DB.Create(&accessToken).Error; err != nil {
		return accessToken, err
	}
	return accessToken, nil
}
