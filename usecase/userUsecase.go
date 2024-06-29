package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sypnasis-golang-test-ecommerce/config"
	"sypnasis-golang-test-ecommerce/models"
	"sypnasis-golang-test-ecommerce/repository"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Claims struct {
	UserID uint64
	jwt.MapClaims
}

type UserBody struct {
	Email    string
	Password string
}

type Response struct {
	Success bool
	Message string
	Data    interface{}
}

// CreateNewUser creates a new user in the database.
func CreateNewUser(DB *gorm.DB, userBody UserBody) (models.User, error) {
	var user models.User
	user, _ = repository.GetUserByEmail(DB, userBody.Email)

	if user.ID != 0 {
		return user, errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userBody.Password), bcrypt.DefaultCost)

	if err != nil {
		return user, err
	}

	user.Email = userBody.Email
	user.Password = string(hash)

	return repository.CreateUser(DB, user)
}

type LoginUserInput struct {
	Email    string
	Password string
}

var ctx = context.Background()

// LoginUser checks the user's credentials and returns a JWT token if successful.
func LoginUser(RDB *redis.Client, DB *gorm.DB, UserInput LoginUserInput) (string, error) {
	user, err := repository.GetUserByEmail(DB, UserInput.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))

	if err != nil {
		return "", errors.New("invalid password")
	}

	uuid := uuid.New()
	uuidStr := uuid.String()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": uuidStr,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	expiredToken := time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(config.ConfigENV.JWTKey))

	if err != nil {
		return "", err
	}

	accessToken := models.AccessToken{
		UserID:    user.ID,
		AccessID:  uuidStr,
		ExpiredAt: expiredToken,
	}

	key := fmt.Sprintf("%d#%s", user.ID, uuidStr)

	user.Password = ""

	out, err := json.Marshal(user)

	if err != nil {
		return "", errors.New("failed to marshal user")
	}

	_, err = RDB.SetNX(ctx, key, out, 60*time.Second).Result()

	if err != nil {
		return "", errors.New("failed to marshal user")
	}

	_, err = repository.CreateAccessToken(DB, accessToken)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// AuthorizeJWT validates the JWT token from the authorization header.
func AuthorizeJWT(RDB *redis.Client, DB *gorm.DB, tokenString string) (models.User, error) {
	var user models.User

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return user, fmt.Errorf("invalid token")
		}

		return []byte(config.ConfigENV.JWTKey), nil
	})

	if err != nil {
		return user, errors.New("failed to parse token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return user, errors.New("token expired")
		}

		// find user
		var accessToken models.AccessToken
		DB.First(&accessToken, "access_id = ?", claims["sub"])

		if accessToken.UserID == 0 {
			return user, errors.New("invalid token")
		}

		var user models.User
		key := fmt.Sprintf("%d#%s", user.ID, accessToken.AccessID)

		val, err := RDB.Get(ctx, key).Result()
		if err != nil {
			DB.First(&user, "id = ?", accessToken.UserID)

			key := fmt.Sprintf("%d#%s", user.ID, accessToken.AccessID)

			user.Password = ""

			out, err := json.Marshal(user)

			if err != nil {
				return user, errors.New("failed to marshal user")
			}

			_, err = RDB.SetNX(ctx, key, out, 60*time.Second).Result()

			if err != nil {
				return user, errors.New("failed to marshal user")
			}
		} else {
			err := json.Unmarshal([]byte(val), &user)
			if err != nil {
				return user, errors.New("failed to unmarshal user")
			}
		}

		if user.ID == 0 {
			return user, errors.New("user not found")
		}
		user.Password = ""
		return user, nil
	}

	return user, errors.New("invalid token")
}
