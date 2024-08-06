package bootstrap

import (
	"task_managemet_api/cmd/task_manager/internal/delivery/http"
	"task_managemet_api/cmd/task_manager/internal/repository"
	"task_managemet_api/cmd/task_manager/internal/routes"
	"task_managemet_api/cmd/task_manager/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRouter(client *mongo.Client) *gin.Engine {
	router := gin.Default()
	taskRepository := repository.NewTaskService(client)
	taskUsecase := usecase.NewTaskUseCase(taskRepository)
	taskHandler := http.NewTaskHandler(taskUsecase)
	routes.RegisterTaskRoutes(router, taskHandler)

	userRepository := repository.NewUserServices(client)
	userUsecase := usecase.NEwUserUSecase(userRepository)
	userHandler := http.NewUserHandler(userUsecase)

	routes.RegisterUserRoutes(router, userHandler)

	return router

}
