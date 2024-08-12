package security

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetFromToken interface {
	GetUserIdFormToken(c *gin.Context) (primitive.ObjectID, error)
	GetEmailFromToken(c *gin.Context) (string, error)
	GetRoleFromToken(c *gin.Context) (string, error)
}

type GetTokenData struct {
}

func (GetTokenData) GetUserIdFormToken(c *gin.Context) (primitive.ObjectID, error) {
	claims, exists := c.Get("claims")
	if !exists {
		return (primitive.ObjectID{}), nil
	}
	userIdStr := claims.(jwt.MapClaims)["userId"].(string)
	userId, errr := primitive.ObjectIDFromHex(userIdStr)
	if errr != nil {
		fmt.Println("error", errr)
		return (primitive.ObjectID{}), errr
	}
	return userId, nil
}

func (GetTokenData) GetEmailFromToken(c *gin.Context) (string, error) {
	claims, exists := c.Get("claims")
	if !exists {
		return "", nil
	}
	email := claims.(jwt.MapClaims)["email"].(string)
	return email, nil
}

func (GetTokenData) GetRoleFromToken(c *gin.Context) (string, error) {
	claims, exists := c.Get("claims")
	if !exists {
		return "", nil
	}
	role := claims.(jwt.MapClaims)["role"].(string)
	return role, nil
}
