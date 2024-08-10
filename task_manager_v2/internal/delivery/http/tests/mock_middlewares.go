package tests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MockAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		mockClaims := jwt.MapClaims{
			"userId": primitive.NewObjectID().Hex(),
			"role":   "user",
			"email":  "test_user@example.com",
			"exp":    1200034567890,
		}
		// Set the mock claims in the context
		c.Set("claims", mockClaims)
		c.Next()
	}
}

func MockISAdminMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Claims not found"})
			c.Abort()
			return
		}

		claims.(map[string]interface{})["role"] = "admin"

		c.Next()
	}
}
