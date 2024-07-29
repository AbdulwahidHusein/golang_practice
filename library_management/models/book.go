package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title  string             `bson:"title" json:"title"`
	Author string             `bson:"author" json:"author"`
	Year   int                `bson:"year" json:"year"`
	Status string             `bson:"status" json:"status"`
}

type BorrowedBook struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId    primitive.ObjectID `bson:"user_id" json:"user_id"`
	BookId    primitive.ObjectID `bson:"book_id" json:"book_id"`
	BorowedAt string             `bson:"borowed_at" json:"borowed_at"`
}
