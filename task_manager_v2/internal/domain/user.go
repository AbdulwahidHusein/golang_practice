package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string             `bson:"email" json:"email" gorm:"unique"`
	Role     string             `bson:"role" json:"role"`
	Password string             `bson:"password" json:"password" gorm:"not null"`
}
