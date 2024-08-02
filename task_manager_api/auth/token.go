package auth

import (
	"fmt"
	"log"
	"task_management_api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userId, role string) (string, error) {
	secretKey := config.GetSecretKey()
	if len(secretKey) == 0 {
		log.Fatal("SECRET_KEY not set in .env file")
	}
	tokenExpiryInt := config.GetTokenExpiry()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"role":   role,
		"exp":    time.Now().Add(time.Hour * time.Duration(tokenExpiryInt)).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	secretKey := config.GetSecretKey()
	if len(secretKey) == 0 {
		return nil, fmt.Errorf("SECRET_KEY not set in .env file")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	// role, ok := claims["role"].(string)
	if !ok {
		return nil, fmt.Errorf("role not found in token claims")
	}

	return claims, nil
}
