package security

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUSerIdFormToken(c *gin.Context) (primitive.ObjectID, error) {
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
