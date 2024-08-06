package bootstrap

import (
	"task_managemet_api/cmd/task_manager/internal/delivery/http"
	"task_managemet_api/cmd/task_manager/internal/repository"
	"task_managemet_api/cmd/task_manager/internal/routes"
	"task_managemet_api/cmd/task_manager/internal/usecase"

	"github.com/gin-gonic/gin"
)

func GetRouter(taskRepo repository.TaskRepository, userRepo repository.UserRepository) *gin.Engine {
	router := gin.Default()

	taskUsecase := usecase.NewTaskUseCase(taskRepo)
	taskHandler := http.NewTaskHandler(taskUsecase)
	routes.RegisterTaskRoutes(router, taskHandler)

	userUsecase := usecase.NEwUserUSecase(userRepo)
	userHandler := http.NewUserHandler(userUsecase)
	routes.RegisterUserRoutes(router, userHandler)

	return router
}
