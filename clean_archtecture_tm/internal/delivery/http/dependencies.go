package http

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetFromToken interface {
	GetUserIdFormToken(c *gin.Context) (primitive.ObjectID, error)
	GetEmailFromToken(c *gin.Context) (string, error)
	GetRoleFromToken(c *gin.Context) (string, error)
}
