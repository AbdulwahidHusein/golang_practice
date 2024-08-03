package config

import (
	"log"
	"os"

	"strconv"

	"github.com/joho/godotenv"
)

func GetSecretKey() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	return os.Getenv("SECRET_KEY")
}

func GetMongoURI() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	MongoURI := os.Getenv("MONGODB_URI")
	return MongoURI
}

func GetTokenExpiry() int {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	tokenExpiry := os.Getenv("TOKEN_EXPIRY")

	tokenExpiryInt, err := strconv.Atoi(tokenExpiry)

	if err != nil {
		tokenExpiryInt = 24 // default to 24 hours if TOKEN_EXPIRY is not set or invalid
	}

	return tokenExpiryInt
}
