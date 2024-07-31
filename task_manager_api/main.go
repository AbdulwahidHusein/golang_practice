package main

//"github.com/gin-gonic/gin"
import (
	"net/http"
	"task_management_api/config"
	"task_management_api/controllers"
	"task_management_api/services"

	"context"
	"log"
	"task_management_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	client, err := config.DBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	router := gin.Default()

	taskservice := services.NewTaskService(client)
	taskController := controllers.NewTaskController(taskservice)

	routes.RegisterTaskRoutes(router, taskController)

	http.ListenAndServe(":8080", router)
}
