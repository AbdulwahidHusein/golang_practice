package routes

import (
	"task_managemet_api/cmd/task_manager/internal/delivery/http"

	"task_managemet_api/cmd/task_manager/internal/middleware"

	"task_managemet_api/cmd/task_manager/pkg/security"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userController *http.UserHandler) {
	r := router.Group("/auth")
	r.POST("/register", userController.AddUser)
	r.POST("/login", userController.LoginUser)
	r.PUT("/user", userController.UpdateUser, middleware.AuthMiddleware())
	r.DELETE("/user", userController.DeleteUser, middleware.AuthMiddleware())
	r.GET("/user", userController.GetUser, middleware.AuthMiddleware())
	r.POST("/refresh", security.RefreshTokenHandler, middleware.AuthMiddleware())
}
