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

func GetTokenExpiry() (int, int) {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	AcessExpiry := os.Getenv("ACCESSES_TOKEN_EXPIRY")
	RefreshExpiry := os.Getenv("REFRESH_TOKEN_EXPIRY")

	accessesExpiryInt, err := strconv.Atoi(AcessExpiry)
	if err != nil {
		accessesExpiryInt = 24 // default to 24 hours if TOKEN_EXPIRY is not set or invalid
	}
	refteshExpiryInt, err := strconv.Atoi(RefreshExpiry)

	if err != nil {
		refteshExpiryInt = 48
	}

	return (accessesExpiryInt), (refteshExpiryInt)
}
