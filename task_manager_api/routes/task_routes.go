package routes

import (
	"task_management_api/controllers"

	"task_management_api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine, taskController *controllers.TaskController) {
	authorized := router.Group("/", middleware.AuthMiddleware())

	authorized.GET("/tasks", taskController.GetTasks)
	authorized.GET("/tasks/:id", taskController.GetTask)
	authorized.PUT("/tasks/:id", taskController.UpdateTask)
	authorized.POST("/tasks", taskController.CreateTask)
	authorized.DELETE("/tasks/:id", taskController.DeleteTask)
	authorized.GET("/tasks/done", taskController.GetDoneTasks)
	authorized.GET("/tasks/undone", taskController.GetUndoneTasks)
}
