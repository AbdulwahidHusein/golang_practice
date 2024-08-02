package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func CreateToken(email, role string) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	if len(secretKey) == 0 {
		log.Fatal("SECRET_KEY not set in .env file")
	}
	tokenExpiry := os.Getenv("TOKEN_EXPIRY")
	tokenExpiryInt, err := strconv.Atoi(tokenExpiry)
	if err != nil {
		tokenExpiryInt = 24 // default to 24 hours if TOKEN_EXPIRY is not set or invalid
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * time.Duration(tokenExpiryInt)).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", fmt.Errorf("Error loading .env file")
	}
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	if len(secretKey) == 0 {
		return "", fmt.Errorf("SECRET_KEY not set in .env file")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", jwt.ErrSignatureInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", fmt.Errorf("role not found in token claims")
	}

	return role, nil
}
