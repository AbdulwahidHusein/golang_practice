package db

import (
	"context"
	"log"
	"task_managemet_api/cmd/task_manager/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo() (*mongo.Client, error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.GetMongoURI()))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil
}
