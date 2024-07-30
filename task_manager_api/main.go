package main

//"github.com/gin-gonic/gin"
import (
	"net/http"
	"task_management_api/controllers"
	"task_management_api/services"

	"context"
	"task_management_api/routes"

	"github.com/go-chi/chi/v5"
)

func main() {
	err, cancel := context.WithCancel(context.Background())

	defer cancel()

	router := chi.NewRouter()
	taskservice := services.NewTaskService()
	taskController := controllers.NewTaskController(taskservice)

	routes.RegisterTaskRoutes(router, taskController)

	http.ListenAndServe(":8080", router)
}
