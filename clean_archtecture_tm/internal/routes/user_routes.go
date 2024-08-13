package routes

import (
	"task_managemet_api/cmd/task_manager/internal/delivery/http"

	"task_managemet_api/cmd/task_manager/internal/middleware"

	"task_managemet_api/cmd/task_manager/pkg/security"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userController *http.UserHandler) {
	r := router.Group("/auth")
	tu := security.TokenUtil{}
	r.POST("/register", userController.AddUser)
	r.POST("/login", userController.LoginUser)
	r.PUT("/user", middleware.AuthMiddleware(), userController.UpdateUser)
	r.DELETE("/user", middleware.AuthMiddleware(), userController.DeleteUser)
	r.GET("/user", middleware.AuthMiddleware(), userController.GetUser)
	r.POST("/refresh", middleware.AuthMiddleware(), tu.RefreshTokenHandler)

	admin := router.Group("/admin", middleware.AuthMiddleware(), middleware.ISAdminMiddleWare())
	admin.PUT("/approve/:id", userController.ApproveUser)
	admin.PUT("/disapprove/:id", userController.DisApproveUser)
	admin.POST("/create-admin", userController.CreateAdmin)

	admin.PUT("/promote/:id", userController.PromoteUser)
	admin.PUT("/demote/:id", userController.DemoteUser)
}
