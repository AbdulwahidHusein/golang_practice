package repository

import (
	"context"
	"task_managemet_api/cmd/task_manager/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	AddUser(user *domain.User) error
	IsEmptyCollection(ctx context.Context) (bool, error)
	DeleteUser(id primitive.ObjectID) error
	UpdateUser(id primitive.ObjectID, user *domain.User) (*domain.User, error)
	GetUSerById(id primitive.ObjectID) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}
