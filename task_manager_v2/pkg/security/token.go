package security

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"task_managemet_api/cmd/task_manager/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(userId, role, email string) (string, string, error) {
	secretKey := config.GetEnvs()["SECRET_KEY"]
	if len(secretKey) == 0 {
		log.Fatal("SECRET_KEY not set in .env file")
	}
	accessExpiry, err1 := strconv.Atoi(config.GetEnvs()["ACCESSES_TOKEN_EXPIRY"])
	RefreshExpiry, err2 := strconv.Atoi(config.GetEnvs()["REFRESH_TOKEN_EXPIRY"])
	if err1 != nil || err2 != nil {
		log.Fatal("TOKEN_EXPIRY not set in .env file")
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"role":   role,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * time.Duration(accessExpiry)).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"role":   role,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * time.Duration(RefreshExpiry)).Unix(),
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	secretKey := config.GetEnvs()["SECRET_KEY"]
	if len(secretKey) == 0 {
		return nil, fmt.Errorf("SECRET_KEY not set in .env file")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
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

func RefreshToken(refreshToken string) (string, string, error) {

	claims, err := VerifyToken(refreshToken)
	if err != nil {
		return "", "", err
	}
	userId, role, email := claims["userId"].(string), claims["role"].(string), claims["email"].(string)
	newAccess, newRefresh, err := CreateToken(userId, role, email)

	if err != nil {
		return "", "", err
	}
	return newAccess, newRefresh, nil
}

func RefreshTokenHandler(c *gin.Context) {
	refreshToken := c.PostForm("refresh_token")
	access, refresh, err := RefreshToken(refreshToken)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": []map[string]string{{"access": access, "refresh": refresh}}})

}
