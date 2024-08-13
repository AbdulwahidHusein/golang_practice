package routes

import (
	"task_managemet_api/cmd/task_manager/internal/delivery/http"
	"task_managemet_api/cmd/task_manager/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine, taskController *http.TaskHandler) {
	authOnly := router.Group("/", middleware.AuthMiddleware())
	authOnly.GET("/tasks", taskController.GetTasks)
	authOnly.GET("/tasks/status/:status", taskController.GetTaskByStatus) // Changed route
	authOnly.GET("/tasks/:id", taskController.GetTask)

	adminOnly := router.Group("/admin", middleware.AuthMiddleware(), middleware.ISAdminMiddleWare())
	adminOnly.GET("/tasks", taskController.GetTasks)
	adminOnly.GET("/tasks/:id", taskController.GetTask)
	adminOnly.PUT("/tasks/:id", taskController.UpdateTask)
	adminOnly.POST("/tasks", taskController.CreateTask)
	adminOnly.DELETE("/tasks/:id", taskController.DeleteTask)
	adminOnly.GET("/tasks/status/:status", taskController.GetTaskByStatus) // Changed route

}
