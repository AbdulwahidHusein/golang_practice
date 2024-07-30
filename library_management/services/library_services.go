package services

import (
	"context"
	"library_management/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LibraryService struct {
	book_collection   *mongo.Collection
	borrow_collection *mongo.Collection
}

func NewLibraryService(client *mongo.Client) *LibraryService {
	book_collection := client.Database("library_management").Collection("books")
	borrow_collection := client.Database("library_management").Collection("borrows")
	return &LibraryService{book_collection, borrow_collection}
}

func (l *LibraryService) AddBook(book *models.Book) error {
	_, err := l.book_collection.InsertOne(context.TODO(), book)
	if err != nil {
		return err
	}
	return nil
}

func (l *LibraryService) GetBooks() ([]*models.Book, error) {
	var books []*models.Book
	cursor, err := l.book_collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var book models.Book
		err := cursor.Decode(&book)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}

func (l *LibraryService) GetBook(id primitive.ObjectID) (*models.Book, error) {
	var book models.Book
	err := l.book_collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (l *LibraryService) UpdateBook(id primitive.ObjectID, book *models.Book) error {
	_, err := l.book_collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": book})
	if err != nil {
		return err
	}
	return nil
}

func (l *LibraryService) DeleteBook(id primitive.ObjectID) error {
	_, err := l.book_collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (l *LibraryService) BorrowBook(borrow *models.BorrowedBook) error {
	bookId := borrow.BookId
	_, err1 := l.book_collection.UpdateOne(context.TODO(), bson.M{"_id": bookId}, bson.M{"$set": bson.M{"status": "borrowed"}})
	if err1 != nil {
		return err1
	}
	_, err := l.borrow_collection.InsertOne(context.TODO(), borrow)
	if err != nil {
		return err
	}
	return nil
}

func (l *LibraryService) GetBorrowedBooks(userId primitive.ObjectID) ([]*models.BorrowedBook, error) {
	var borrowedBooks []*models.BorrowedBook
	cursor, err := l.borrow_collection.Find(context.TODO(), bson.M{"user_id": userId})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var borrowedBook models.BorrowedBook
		err := cursor.Decode(&borrowedBook)
		if err != nil {
			return nil, err
		}
		borrowedBooks = append(borrowedBooks, &borrowedBook)
	}
	return borrowedBooks, nil
}

func (l *LibraryService) UnborrowBook(borrow *models.BorrowedBook) error {
	bookId := borrow.BookId
	_, err1 := l.book_collection.UpdateOne(context.TODO(), bson.M{"_id": bookId}, bson.M{"$set": bson.M{"status": "available"}})
	if err1 != nil {
		return err1
	}
	_, err := l.borrow_collection.DeleteOne(context.TODO(), bson.M{"_id": borrow.ID})
	if err != nil {
		return err
	}
	return nil
}

func (l *LibraryService) GetAvailableBooks() ([]*models.Book, error) {
	var books []*models.Book
	cursor, err := l.book_collection.Find(context.TODO(), bson.M{"status": "available"})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var book models.Book
		err := cursor.Decode(&book)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}
