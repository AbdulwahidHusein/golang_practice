package main

import (
	"go_jwt/config"
	"go_jwt/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.Initialize(); err != nil {
		log.Fatal(err)
	}
	client := config.Client
	service := models.NewUserService(client)

	router := gin.Default()
	router.POST("/register", service.CreateUser)
	router.POST("/login", service.LoginUser)
	router.GET("/users", service.GetUsers)

	log.Fatal(router.Run(":8080"))
}
