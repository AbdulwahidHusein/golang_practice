package config

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Initialize() error {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	DB_URL := os.Getenv("MONGODB_URI")
	if DB_URL == "" {
		log.Fatal("MONGO_URL environment variable not set")
		return errors.New("MONGO_URL environment variable not set")
	}

	var err error
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(DB_URL))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
		return err
	}

	// Check the connection
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
		return err
	}

	log.Println("Connected to MongoDB!")
	return nil
}
