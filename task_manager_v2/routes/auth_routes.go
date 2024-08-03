package routes

import (
	"task_management_api/controllers"

	"task_management_api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userController *controllers.UserController) {
	r := router.Group("/auth")
	r.POST("/register", userController.AddUser)
	r.POST("/login", userController.LoginUser)
	r.PUT("/user", userController.UpdateUser, middleware.AuthMiddleware())
	r.DELETE("/user", userController.DeleteUser, middleware.AuthMiddleware())
	r.GET("/user", userController.GetUser, middleware.AuthMiddleware())
}
