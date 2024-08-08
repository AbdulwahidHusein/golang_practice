package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ISAdminMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Claims not found"})
			c.Abort()
			return
		}
		role, ok := claims.(jwt.MapClaims)["role"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found"})
			c.Abort()
			return
		}
		if !ok || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admins only you are ", "role": role})
			c.Abort()
			return
		}
		c.Next()
	}
}
