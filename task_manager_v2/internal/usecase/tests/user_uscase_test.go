package tests

import (
	"context"
	"testing"

	"task_managemet_api/cmd/task_manager/internal/domain"
	"task_managemet_api/cmd/task_manager/internal/usecase"

	// "task_managemet_api/cmd/task_manager/pkg/security"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mock for UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) AddUser(user *domain.User) (*domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) IsEmptyCollection(ctx context.Context) (bool, error) {
	args := m.Called(ctx)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepository) DeleteUser(id primitive.ObjectID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateUser(id primitive.ObjectID, user *domain.User) (*domain.User, error) {
	args := m.Called(id, user)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetUSerById(id primitive.ObjectID) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

// Test for AddUser method
// Test for AddUser method
func TestUserUsecase_AddUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUsecase := usecase.NEwUserUSecase(mockRepo)

	user := &domain.User{
		Email:    "testAbc@gmail.com",
		Password: "Stron@#$adgP@ssw0rd123",
	}

	// Mock the GetUserByEmail to return nil (indicating the user does not exist)
	mockRepo.On("GetUserByEmail", mock.Anything).Return(nil, nil)

	// Mock the IsEmptyCollection to return false
	mockRepo.On("IsEmptyCollection", mock.Anything).Return(false, nil)

	// Mock the AddUser to return the user
	mockRepo.On("AddUser", user).Return(user, nil)

	createdUser, err := userUsecase.AddUser(user)

	if err != nil {
		t.Fatalf("expected no error while adding user, but got: %v", err)
	}
	if createdUser == nil {
		t.Fatal("expected non-nil user, but got nil")
	}
	if createdUser.Role != "user" {
		t.Errorf("expected user role to be 'user', but got: %v", createdUser.Role)
	}

	mockRepo.AssertExpectations(t)
}

func TestAddUserInvalidPassword(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUsecase := usecase.NEwUserUSecase(mockRepo)

	user := &domain.User{
		Email:    "testAbc@gmail.com",
		Password: "password",
	}
	mockRepo.On("GetUserByEmail", mock.Anything).Return(nil, nil)
	mockRepo.On("IsEmptyCollection", mock.Anything).Return(false, nil)
	mockRepo.On("AddUser", user).Return(user, nil)

	_, err := userUsecase.AddUser(user)
	require.Error(t, err)

	mockRepo.AssertExpectations(t)
}

func TestFirstUserAdmin(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUsecase := usecase.NEwUserUSecase(mockRepo)

	user := &domain.User{
		Email:    "testAbc@gmail.com",
		Password: "passwordA@#$123",
	}
	mockRepo.On("GetUserByEmail", mock.Anything).Return(nil, nil)
	mockRepo.On("IsEmptyCollection", mock.Anything).Return(true, nil)
	mockRepo.On("AddUser", user).Return(user, nil)
	createdUser, err := userUsecase.AddUser(user)

	if err != nil {
		t.Fatalf("expected no error while adding user, but got: %v", err)
	}
	require.Equal(t, "admin", createdUser.Role)

	mockRepo.AssertExpectations(t)
}

// Test for DeleteUser method
func TestUserUsecase_DeleteUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUsecase := usecase.NEwUserUSecase(mockRepo)

	userID := primitive.NewObjectID()

	// Test unauthorized deletion
	err := userUsecase.DeleteUser(userID, primitive.NewObjectID())
	assert.EqualError(t, err, "unauthorized deletion", "expected unauthorized deletion error, but got: %v", err)

	// Test authorized deletion
	mockRepo.On("DeleteUser", userID).Return(nil)
	err = userUsecase.DeleteUser(userID, userID)
	assert.NoError(t, err, "expected no error while deleting user, but got: %v", err)
	mockRepo.AssertExpectations(t)
}

// Test for LoginUser method
func TestUserUsecase_LoginUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUsecase := usecase.NEwUserUSecase(mockRepo)

	user := &domain.User{
		Email:    "test@example.com",
		Password: "$2a$12$ViW2yO/fbVtIbHDmPIgjNOEj6QqJqgWen33FFAhFT.0UCQhNDs1Ny", // Simulate a hashed password
		Role:     "user",
	}

	// Mock the GetUserByEmail to return the user
	mockRepo.On("GetUserByEmail", "test@example.com").Return(user, nil)

	// Mock the password comparison
	// security.ComparePassword = func(hashedPassword, password string) error {
	// 	return nil
	// }

	// // Mock the token creation
	// security.CreateToken = func(userID, role, email string) (string, string, error) {
	// 	return "access-token", "refresh-token", nil
	// }

	accessToken, refreshToken, err := userUsecase.LoginUser("test@example.com", "StA234!@#rongP@ssw0rd")

	assert.NoError(t, err, "expected no error while logging in, but got: %v", err)
	assert.NotNil(t, accessToken, "expected non-nil access token, but got nil")
	assert.NotNil(t, refreshToken, "expected non-nil refresh token, but got nil")
	mockRepo.AssertExpectations(t)
}

func TestInvalidPasswordLogin(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUsecase := usecase.NEwUserUSecase(mockRepo)

	user := &domain.User{
		Email:    "test@example.com",
		Password: "$2a$12$ViW2yO/fbVtIbHDmPIgjNOEj6QqJqgWen33FFAhFT.0UCQhNDs1N", //  a hashed password
		Role:     "user",
	}
	mockRepo.On("GetUserByEmail", "test@example.com").Return(user, nil)
	accessToken, refreshToken, err := userUsecase.LoginUser("test@example.com", "StA234!@#rongP@ssw0rd")

	assert.Error(t, err, "expected error while logging in since the hashed password is of incorrect, but got: %v", err)
	assert.Nil(t, accessToken, "expected nil access token, but got: %v", accessToken)
	assert.Nil(t, refreshToken, "expected nil refresh token, but got: %v", refreshToken)
	mockRepo.AssertExpectations(t)
}
