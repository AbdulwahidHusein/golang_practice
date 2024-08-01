package controllers

import (
	"encoding/json"
	"library_management/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LibraryManager interface {
	GetBookService(id primitive.ObjectID) (*models.Book, error)
	GetBooksService() ([]*models.Book, error)
	AddBookService(book *models.Book) error
	UpdateBookService(id primitive.ObjectID, book *models.Book) error
	DeleteBookService(id primitive.ObjectID) error
	BorrowBookService(borrow *models.BorrowedBook) error
	GetBorrowedBooksService(userId primitive.ObjectID) ([]*models.BorrowedBook, error)
	ReturnBookService(borrow *models.BorrowedBook) error
	GetAvailableBooksService() ([]*models.Book, error)
}

type LibraryController struct {
	LibraryService LibraryManager
}

func NewLibraryController(service LibraryManager) *LibraryController {
	return &LibraryController{service}
}

func (l *LibraryController) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := l.LibraryService.GetBooksService()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func (l *LibraryController) GetBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bookId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book, err := l.LibraryService.GetBookService(bookId)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (l *LibraryController) GetAvailableBooks(w http.ResponseWriter, r *http.Request) {
	books, err := l.LibraryService.GetAvailableBooksService()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func (l *LibraryController) AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err = l.LibraryService.AddBookService(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (l *LibraryController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err = l.LibraryService.UpdateBookService(book.ID, &book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (l *LibraryController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bookId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	err = l.LibraryService.DeleteBookService(bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (l *LibraryController) BorrowBook(w http.ResponseWriter, r *http.Request) {
	var borrow models.BorrowedBook
	err := json.NewDecoder(r.Body).Decode(&borrow)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err = l.LibraryService.BorrowBookService(&borrow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(borrow)
}

func (l *LibraryController) ReturnBook(w http.ResponseWriter, r *http.Request) {
	var borrow models.BorrowedBook
	err := json.NewDecoder(r.Body).Decode(&borrow)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err = l.LibraryService.ReturnBookService(&borrow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(borrow)
}

func (l *LibraryController) GetUserBorrowedBooks(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	borrowedBooks, err := l.LibraryService.GetBorrowedBooksService(userId)
	if err != nil {
		http.Error(w, "Error retrieving borrowed books", http.StatusInternalServerError)
		return
	}
	if len(borrowedBooks) == 0 {
		http.Error(w, "No borrowed books found for this user", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(borrowedBooks)
}
