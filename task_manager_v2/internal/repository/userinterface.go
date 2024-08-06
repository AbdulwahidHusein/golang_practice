package repository

import (
	"task_managemet_api/cmd/task_manager/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	AddUser(user *domain.User) error
	DeleteUser(id primitive.ObjectID) error
	UpdateUser(id primitive.ObjectID, user *domain.User) error
	GetUser(id primitive.ObjectID) (*domain.User, error)
	LoginUser(email string) (*domain.User, error)
	CheckUser(email string) (*domain.User, error)
}
