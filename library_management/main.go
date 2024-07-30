package main

import (
	"context"
	"library_management/config"
	"library_management/controllers"
	"library_management/routes"
	"library_management/services"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// connect to db
	client, err := config.DBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	BookService := services.NewLibraryService(client)
	MemberService := services.NewUserService(client)

	BookController := controllers.NewLibraryController(BookService)
	MemberController := controllers.NewUserController(MemberService)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	routes.RegisterBookRouter(router, BookController)
	routes.RegisterUserRouter(router, MemberController)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
