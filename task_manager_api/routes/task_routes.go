package routes

import (
	"task_management_api/controllers"

	"github.com/go-chi/chi/v5"
)

func RegisterTaskRoutes(router *chi.Mux, taskController *controllers.TaskController) {
	router.Get("/tasks", taskController.GetTasks)
	router.Get("/tasks/{id}", taskController.GetTask)
	router.Put("/tasks/{id}", taskController.UpdateTask)
	router.Post("/tasks", taskController.CreateTask)
}
