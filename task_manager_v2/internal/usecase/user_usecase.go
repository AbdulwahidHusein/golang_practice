package usecase

import (
	"context"
	"task_managemet_api/cmd/task_manager/internal/domain"
	"task_managemet_api/cmd/task_manager/pkg/security"
	"time"

	"task_managemet_api/cmd/task_manager/internal/repository"

	"task_managemet_api/cmd/task_manager/pkg/validation"

	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUseCaseInterface interface {
	AddUser(user *domain.User) (*domain.User, error)
	CreateAdmin(admin *domain.User) (*domain.User, error)
	DeleteUser(deleterID primitive.ObjectID, tobeDeletedID primitive.ObjectID) error
	UpdateUser(id primitive.ObjectID, user *domain.User) *domain.User
	GetUser(id primitive.ObjectID) (*domain.User, error)
	LoginUser(email string, password string) (string, string, error)
	ActivateUser(id primitive.ObjectID) (*domain.User, error)
	DeactivateUser(id primitive.ObjectID) (*domain.User, error)
}

type UserUsecase struct {
	userRepository repository.UserRepository
}

func NEwUserUSecase(userRepository repository.UserRepository) UserUsecase {
	return UserUsecase{
		userRepository: userRepository,
	}
}

func (u UserUsecase) AddUser(user *domain.User) (*domain.User, error) {
	if !validation.IsValidEmail(user.Email) {
		return nil, errors.New("invalid email format")
	}
	if !validation.IsValidPassword(user.Password) {
		return nil, errors.New("password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character")
	}
	usr, err := u.userRepository.GetUserByEmail(user.Email)
	if usr != nil {
		return nil, errors.New("user with this email already exists")
	}
	if err != nil {
		return nil, err
	}

	hashedPassword, _ := security.EncryptPassword(user.Password)
	user.Password = hashedPassword
	user.Isactivated = false

	if isdbemp, _ := u.userRepository.IsEmptyCollection(context.Background()); isdbemp {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}
	user.CreatedAt = time.Now()

	createdUser, err := u.userRepository.AddUser(user)
	if err != nil {
		return nil, err
	}
	if createdUser == nil {
		return nil, errors.New("user not created")
	}
	return createdUser, nil

}

// func (u UserUsecase) AddAdmin(ctx context.Context) (bool, error) {

// }

func (u UserUsecase) DeleteUser(deleterID primitive.ObjectID, tobeDeletedID primitive.ObjectID) error {
	if deleterID != tobeDeletedID {
		return errors.New("unauthorized deletion")
	}
	return u.userRepository.DeleteUser(deleterID)
}

func (u UserUsecase) UpdateUser(id primitive.ObjectID, user *domain.User) *domain.User {
	DbUser, err := u.userRepository.GetUSerById(id)
	if err != nil {
		return nil
	}
	user.ID = id
	user.Role = DbUser.Role
	user.Isactivated = DbUser.Isactivated
	user.CreatedAt = DbUser.CreatedAt
	user.Email = DbUser.Email
	user.Password = DbUser.Password
	user, err1 := u.userRepository.UpdateUser(id, user)
	if err1 != nil {
		return nil
	}
	return user
}

func (u UserUsecase) GetUser(id primitive.ObjectID) (*domain.User, error) {
	return u.userRepository.GetUSerById(id)
}

func (u UserUsecase) LoginUser(email string, password string) (string, string, error) {
	realUser, err := u.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", "", err
	}
	hashedPassword := realUser.Password
	if err := security.ComparePassword(hashedPassword, password); err != nil {
		return "", "", err
	}
	accessTokenString, refreshTokenString, err := security.CreateToken(realUser.ID.Hex(), realUser.Role, realUser.Email)
	if err != nil {
		return "", "", err
	}
	return accessTokenString, refreshTokenString, nil

}

func (u UserUsecase) DeactivateUser(userID primitive.ObjectID) (*domain.User, error) {
	user, err := u.userRepository.GetUSerById(userID)
	if err != nil {
		return nil, err
	}
	user.Isactivated = false
	usr, err1 := u.userRepository.UpdateUser(user.ID, user)
	if err1 != nil {
		return nil, err1
	}
	return usr, nil
}

func (u UserUsecase) ActivateUser(id primitive.ObjectID) (*domain.User, error) {
	user, err := u.userRepository.GetUSerById(id)
	if err != nil {
		return nil, err
	}
	user.Isactivated = true

	usr, err1 := u.userRepository.UpdateUser(user.ID, user)
	if err1 != nil {
		return nil, err1
	}
	return usr, nil
}

func (u UserUsecase) CreateAdmin(user *domain.User) (*domain.User, error) {
	if !validation.IsValidEmail(user.Email) {
		return nil, errors.New("invalid email format")
	}
	if !validation.IsValidPassword(user.Password) {
		return nil, errors.New("password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character")
	}
	usr, err := u.userRepository.GetUserByEmail(user.Email)
	if usr != nil {
		return nil, errors.New("user with this email already exists")
	}
	if err != nil {
		return nil, err
	}

	hashedPassword, _ := security.EncryptPassword(user.Password)
	user.Password = hashedPassword
	user.Isactivated = true
	user.Role = "admin"
	return u.userRepository.AddUser(user)
}
