package usecase

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

type UserUsecase struct {
	userRepository UserRepository
}

func NEwUserUSecase(userRepository UserRepository) UserUsecase {
	return UserUsecase{
		userRepository: userRepository,
	}
}

func (u UserUsecase) AddUser(user *domain.User) error {
	return u.userRepository.AddUser(user)
}

func (u UserUsecase) DeleteUser(id primitive.ObjectID) error {
	return u.userRepository.DeleteUser(id)
}

func (u UserUsecase) UpdateUser(id primitive.ObjectID, user *domain.User) error {
	return u.userRepository.UpdateUser(id, user)
}

func (u UserUsecase) GetUser(id primitive.ObjectID) (*domain.User, error) {
	return u.userRepository.GetUser(id)
}

func (u UserUsecase) LoginUser(email string) (*domain.User, error) {
	return u.userRepository.LoginUser(email)
}

func (u UserUsecase) CheckUser(email string) (*domain.User, error) {
	return u.userRepository.CheckUser(email)
}
