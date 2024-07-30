package routes

import (
	"library_management/controllers"

	"github.com/go-chi/chi/v5"
)

func RegisterBookRouter(router *chi.Mux, BookController *controllers.LibraryController) {
	router.Post("/books", BookController.AddBook)
	router.Get("/books", BookController.GetBooks)
	router.Get("/books/{id}", BookController.GetBook)
	router.Put("/books/{id}", BookController.UpdateBook)
	router.Delete("/books/{id}", BookController.DeleteBook)

	router.Post("/borrow", BookController.BorrowBook)
	router.Post("/return", BookController.ReturnBook)
	router.Get("/available", BookController.GetAvailableBooks)
	router.Get("/borrowed/{id}", BookController.GetUserBorrowedBooks)

}

func RegisterUserRouter(router *chi.Mux, UserController *controllers.UserController) {
	router.Post("/users", UserController.AddUser)
	router.Get("/users", UserController.GetUsers)
	router.Get("/users/{id}", UserController.GetUser)
	router.Put("/users/{id}", UserController.UpdateUser)
	router.Delete("/users/{id}", UserController.DeleteUser)
}
