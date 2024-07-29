package controllers

import (
	"encoding/json"
	"library_management/models"
	"library_management/services"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LibraryController struct {
	LibraryService *services.LibraryService
}

func NewLibraryController(service *services.LibraryService) *LibraryController {
	return &LibraryController{service}
}

func (l *LibraryController) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := l.LibraryService.GetBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

func (l *LibraryController) GetBook(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	bookId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	book, err := l.LibraryService.GetBook(bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (l *LibraryController) GetAvailableBooks(w http.ResponseWriter, r *http.Request) {
	books, err := l.LibraryService.GetAvailableBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

func (l *LibraryController) AddBook(w http.ResponseWriter, r *http.Request) {

	var book models.Book
	err1 := json.NewDecoder(r.Body).Decode(&book)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}

	err2 := l.LibraryService.AddBook(&book)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
}

func (l *LibraryController) UpdateBook(w http.ResponseWriter, r *http.Request) {

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (l *LibraryController) DeleteBook(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	bookId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err2 := l.LibraryService.DeleteBook(bookId)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}

}

func (l *LibraryController) BorrowBook(w http.ResponseWriter, r *http.Request) {

	var borrow models.BorrowedBook
	err := json.NewDecoder(r.Body).Decode(&borrow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (l *LibraryController) ReturnBook(w http.ResponseWriter, r *http.Request) {

	var borrow models.BorrowedBook
	err := json.NewDecoder(r.Body).Decode(&borrow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
