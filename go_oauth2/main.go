package main

import (
	"github.com/Siddheshk02/go-oauth2/config"
	"github.com/Siddheshk02/go-oauth2/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.GoogleConfig()

	app.Get("/google_login", controllers.GoogleLogin)
	app.Post("/google_callback", controllers.GoogleCallback)

	app.Listen(":8080")

}
