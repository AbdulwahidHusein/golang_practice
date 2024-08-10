package db

import (
	"context"
	"fmt"
	"log"
	"task_managemet_api/cmd/task_manager/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo() (*mongo.Client, error) {
	fmt.Println("mongo uri is", config.GetEnvs()["MONGODB_URI"])
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.GetEnvs()["MONGODB_URI"]))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil
}
