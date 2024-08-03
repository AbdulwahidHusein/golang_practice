package routes

import (
	"task_management_api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userController *controllers.UserController) {
	r := router.Group("/auth")
	r.POST("/register", userController.AddUser)
	r.POST("/login", userController.LoginUser)
	r.PUT("/user", userController.UpdateUser)
	r.DELETE("/user", userController.DeleteUser)
}
