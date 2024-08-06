package routes

import (
	"task_managemet_api/cmd/task_manager/internal/delivery/http"

	"task_managemet_api/cmd/task_manager/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine, taskController *http.TaskHandler) {
	taskroute := router.Group("/", middleware.AuthMiddleware())

	taskroute.GET("/tasks", taskController.GetTasks)
	taskroute.GET("/tasks/:id", taskController.GetTask)
	taskroute.PUT("/tasks/:id", taskController.UpdateTask)
	taskroute.POST("/tasks", taskController.CreateTask)
	taskroute.DELETE("/tasks/:id", taskController.DeleteTask)
	taskroute.GET("/tasks/done", taskController.GetDoneTasks)
	taskroute.GET("/tasks/undone", taskController.GetUndoneTasks)
}
