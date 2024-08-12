package middleware

import (
	"net/http"

	"strings"
	"task_managemet_api/cmd/task_manager/pkg/security"

	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			c.Abort()
			return
		}
		authParts := strings.Split(authHeader, " ")

		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}
		tu := security.TokenUtil{}
		claims, err := tu.VerifyToken(authParts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		fmt.Println(claims)
		c.Next()
	}
}
