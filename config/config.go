package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigStruct struct {
	RedisAddr     string
	DBUrl         string
	JWTKey        string
	RedisPassword string
}

var ConfigENV *ConfigStruct

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ConfigENV = &ConfigStruct{
		RedisAddr:     os.Getenv("REDIS_ADDR"),
		DBUrl:         os.Getenv("DATABASE_URL"),
		JWTKey:        os.Getenv("JWT_KEY"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
	}

}
