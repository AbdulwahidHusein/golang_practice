package routes

import (
	"task_management_api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine, taskController *controllers.TaskController) {
	router.GET("/tasks", taskController.GetTasks)
	router.GET("/tasks/:id", taskController.GetTask)
	router.PUT("/tasks/:id", taskController.UpdateTask)
	router.POST("/tasks", taskController.CreateTask)
	router.DELETE("/tasks/:id", taskController.DeleteTask)
}
