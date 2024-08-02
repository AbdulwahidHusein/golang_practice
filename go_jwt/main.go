package main

import (
	"go_jwt/auth"
	"go_jwt/config"
	"log"

	"go_jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.Initialize(); err != nil {
		log.Fatal(err)
	}
	client := config.Client
	service := auth.NewUserService(client)

	router := gin.Default()
	router.POST("/register", service.CreateUser)
	router.POST("/login", service.LoginUser)
	router.GET("/users", service.GetUsers)
	router.GET("/verify", middlewares.AuthMiddleware(), func(c *gin.Context) { c.JSON(200, gin.H{"data": "valid token"}) })

	log.Fatal(router.Run(":8080"))
}
