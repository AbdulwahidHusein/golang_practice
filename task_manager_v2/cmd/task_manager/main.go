package main

import (
	"net/http"
	"task_managemet_api/cmd/task_manager/cmd/bootstrap"
	"task_managemet_api/cmd/task_manager/internal/repository/mongo"
	"task_managemet_api/cmd/task_manager/pkg/db"

	"context"
	"log"
)

func main() {
	client, err := db.ConnectToMongo()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	taskCollection := client.Database("task_manager_db").Collection("tasks")
	taskRepo := mongo.NewMongoTaskRepository(taskCollection)
	userCollection := client.Database("task_manager_db").Collection("users")
	userRepo := mongo.NewMongoUserRepository(userCollection)

	router := bootstrap.GetRouter(taskRepo, userRepo)

	http.ListenAndServe(":8080", router)
}
