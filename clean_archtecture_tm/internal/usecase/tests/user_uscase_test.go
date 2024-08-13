package tests

import (
	"context"
	"errors"
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

type MockPasswordHasher struct {
	mock.Mock
}

func (m *MockPasswordHasher) ComparePassword(hashedPassword, password string) error {
	args := m.Called(hashedPassword, password)
	return args.Error(0)
}

func (m *MockPasswordHasher) EncryptPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

type MockTokenGenerator struct {
	mock.Mock
}

func (m *MockTokenGenerator) CreateToken(userID, role, email string) (string, string, error) {
	args := m.Called(userID, role, email)
	return args.String(0), args.String(1), args.Error(2)
}

type MockValidator struct {
	mock.Mock
}

func (m *MockValidator) IsValidEmail(email string) bool {
	args := m.Called(email)
	return args.Bool(0)
}

func (m *MockValidator) IsValidPassword(password string) bool {
	args := m.Called(password)
	return args.Bool(0)
}

// Test for AddUser method
// Test for AddUser method
func TestUserUsecase_AddUserPositive(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockHasher := new(MockPasswordHasher)
	mockTokenGen := new(MockTokenGenerator)
	MockValidator := new(MockValidator)

	userUsecase := usecase.NewUserUsecase(mockRepo, mockHasher, mockTokenGen, MockValidator)

	user := &domain.User{
		Email:    "testAbc@gmail.com",
		Password: "Stron@#$adgP@ssw0rd123",
	}

	// Mock the GetUserByEmail to return nil (indicating the user does not exist)
	mockRepo.On("GetUserByEmail", mock.Anything).Return(nil, nil)

	// Mock the IsEmptyCollection to return false
	mockRepo.On("IsEmptyCollection", mock.Anything).Return(false, nil)

	mockHasher.On("EncryptPassword", mock.Anything).Return("Stron@#$adgP@ssw0rd123", nil)
	// Mock the AddUser to return the user
	mockRepo.On("AddUser", user).Return(user, nil)
	MockValidator.On("IsValidEmail", mock.Anything).Return(true)
	MockValidator.On("IsValidPassword", mock.Anything).Return(true)

	createdUser, err := userUsecase.AddUser(user)

	require.NoError(t, err)
	require.Equal(t, user.Email, createdUser.Email, "expected non-empty user email")
	require.Equal(t, user.Role, createdUser.Role, "expected non-empty user role")

	mockRepo.AssertExpectations(t)
}

func TestUserUsecase_AddUserRoleAssignedAsAdmin(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockHasher := new(MockPasswordHasher)
	mockTokenGen := new(MockTokenGenerator)
	MockValidator := new(MockValidator)

	userUsecase := usecase.NewUserUsecase(mockRepo, mockHasher, mockTokenGen, MockValidator)

	user := &domain.User{
		Email:    "testAbc@gmail.com",
		Password: "Stron@#$adgP@ssw0rd123",
		Role:     "admin",
	}

	mockRepo.On("GetUserByEmail", mock.Anything).Return(nil, nil)

	mockRepo.On("IsEmptyCollection", mock.Anything).Return(false, nil)

	mockHasher.On("EncryptPassword", mock.Anything).Return("Stron@#$adgP@ssw0rd123", nil)
	// Mock the AddUser to return the user
	mockRepo.On("AddUser", user).Return(user, nil)
	MockValidator.On("IsValidEmail", mock.Anything).Return(true)
	MockValidator.On("IsValidPassword", mock.Anything).Return(true)

	createdUser, err := userUsecase.AddUser(user)

	require.NoError(t, err)
	require.Equal(t, user.Email, createdUser.Email, "expected non-empty user email")
	require.Equal(t, "user", createdUser.Role, "role must be user role but go %v", createdUser.Role)

	mockRepo.AssertExpectations(t)
}

func TestUserUsecase_AddUserUserAlreadyExist(t *testing.T) {

	mockRepo := new(MockUserRepository)
	mockHasher := new(MockPasswordHasher)
	mockTokenGen := new(MockTokenGenerator)
	MockValidator := new(MockValidator)

	userUsecase := usecase.NewUserUsecase(mockRepo, mockHasher, mockTokenGen, MockValidator)

	user := &domain.User{
		Email:    "testAbc@gmail.com",
		Password: "Stron@#$adgP@ssw0rd123",
	}

	mockRepo.On("GetUserByEmail", mock.Anything).Return(user, nil)
	MockValidator.On("IsValidEmail", mock.Anything).Return(true)
	MockValidator.On("IsValidPassword", mock.Anything).Return(true)

	_, err := userUsecase.AddUser(user)

	require.EqualError(t, err, "user with this email already exists")

}

func TestFirstUserAdmin(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockHasher := new(MockPasswordHasher)
	mockTokenGen := new(MockTokenGenerator)
	MockValidator := new(MockValidator)

	userUsecase := usecase.NewUserUsecase(mockRepo, mockHasher, mockTokenGen, MockValidator)

	user := &domain.User{
		Email:    "testAbc@gmail.com",
		Password: "passwordA@#$123",
	}
	mockRepo.On("GetUserByEmail", mock.Anything).Return(nil, nil)
	mockRepo.On("IsEmptyCollection", mock.Anything).Return(true, nil)
	mockRepo.On("AddUser", user).Return(user, nil)
	mockHasher.On("EncryptPassword", mock.Anything).Return("passwordA@#$123", nil)
	MockValidator.On("IsValidEmail", mock.Anything).Return(true)
	MockValidator.On("IsValidPassword", mock.Anything).Return(true)

	createdUser, err := userUsecase.AddUser(user)
	if err != nil {
		t.Fatalf("expected no error while adding user, but got: %v", err)
	}
	require.Equal(t, "admin", createdUser.Role)

	mockRepo.AssertExpectations(t)
}

// Test for DeleteUser method
func TestUserUsecase_DeleteUserDtScenarios(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockHasher := new(MockPasswordHasher)
	mockTokenGen := new(MockTokenGenerator)
	MockValidator := new(MockValidator)

	userUsecase := usecase.NewUserUsecase(mockRepo, mockHasher, mockTokenGen, MockValidator)

	userID := primitive.NewObjectID()
	mockHasher.On("ComparePassword", mock.Anything, mock.Anything).Return(nil)

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
func TestUserUsecase_LoginUserPositive(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockHasher := new(MockPasswordHasher)
	mockTokenGen := new(MockTokenGenerator)
	MockValidator := new(MockValidator)

	userUsecase := usecase.NewUserUsecase(mockRepo, mockHasher, mockTokenGen, MockValidator)

	user := &domain.User{
		Email:    "test@example.com",
		Password: "$2a$12$ViW2yO/fbVtIbHDmPIgjNOEj6QqJqgWen33FFAhFT.0UCQhNDs1Ny", // Simulate a hashed password
		Role:     "user",
	}

	// Mock the GetUserByEmail to return the user
	mockTokenGen.On("CreateToken", mock.Anything, mock.Anything, mock.Anything).Return("access_token", "refresh_token", nil)
	mockHasher.On("ComparePassword", mock.Anything, mock.Anything).Return(nil)
	mockRepo.On("GetUserByEmail", "test@example.com").Return(user, nil)
	accessToken, refreshToken, err := userUsecase.LoginUser("test@example.com", "StA234!@#rongP@ssw0rd")

	assert.NoError(t, err, "expected no error while logging in, but got: %v", err)
	assert.NotNil(t, accessToken, "expected non-nil access token, but got nil")
	assert.NotNil(t, refreshToken, "expected non-nil refresh token, but got nil")
	mockRepo.AssertExpectations(t)
}

func TestInvalidPasswordLogin(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockHasher := new(MockPasswordHasher)
	mockTokenGen := new(MockTokenGenerator)
	MockValidator := new(MockValidator)

	userUsecase := usecase.NewUserUsecase(mockRepo, mockHasher, mockTokenGen, MockValidator)

	user := &domain.User{
		Email:    "test@example.com",
		Password: "$2a$12$ViW2yO/fbVtIbHDmPIgjNOEj6QqJqgWen33FFAhFT.0UCQhNDs1N", //  a hashed passwor)
		Role:     "user",
	}

	mockHasher.On("ComparePassword", mock.Anything, mock.Anything).Return(errors.New("invalid password"))
	mockRepo.On("GetUserByEmail", "test@example.com").Return(user, nil)

	accessToken, refreshToken, err := userUsecase.LoginUser("test@example.com", "StA234!@#rongP@ssw0rd")

	assert.Error(t, err, "since the hash is not correct it must return an error:")
	assert.Equal(t, "", accessToken, "expected non-nil access token, but got nil")
	assert.Equal(t, "", refreshToken, "expected non-nil refresh token, but got nil")
	mockRepo.AssertExpectations(t)
}

func TestUpdateUserDTScenarios(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockHasher := new(MockPasswordHasher)
	mockTokenGen := new(MockTokenGenerator)
	MockValidator := new(MockValidator)

	userUsecase := usecase.NewUserUsecase(mockRepo, mockHasher, mockTokenGen, MockValidator)

	user := &domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "test@example.com",
		Password: "$2a$12$ViW2yO/fbVtIbHDmPIgjNOEj6QqJqgWen33FFAhFT.0UCQhNDs1Ny", //a hashed password
		Role:     "user",
	}
	UpdatedUser := &domain.User{
		ID:        primitive.NewObjectID(),
		Email:     "test@example.com",
		Password:  "$2a$12$ViW2yO/fbVtIbHDmPIgjNOEj6QqJqgWen33FFAhFT.0UCQhNDs1Ny", //a hashed password
		Role:      "user",
		FirstName: "abdulwahid",
		LastName:  "hs",
	}
	mockRepo.On("UpdateUser", user.ID, UpdatedUser).Return(UpdatedUser, nil)
	mockRepo.On("GetUSerById", user.ID).Return(user, nil)
	usr, err := userUsecase.UpdateUser(user.ID, UpdatedUser)

	assert.NoError(t, err)
	assert.Equal(t, UpdatedUser.Email, user.Email, "expected updated user email to be: %v, but got: %v", UpdatedUser.Email, usr.Email)
	assert.Equal(t, UpdatedUser.Role, user.Role, "expected updated user role to be: %v, but got: %v", UpdatedUser.Role, usr.Role)
	assert.Equal(t, UpdatedUser.Password, user.Password, "password should't change during update")
	assert.Equal(t, UpdatedUser.FirstName, "abdulwahid", "expected updated user first name to be: %v, but got: %v", UpdatedUser.FirstName, usr.FirstName)
	assert.Equal(t, UpdatedUser.LastName, "hs", "expected updated user last name to be: %v, but got: %v", UpdatedUser.LastName, usr.LastName)

	mockRepo.AssertExpectations(t)

}

func TestCreateAdminPositive(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockHasher := new(MockPasswordHasher)
	mockTokenGen := new(MockTokenGenerator)
	MockValidator := new(MockValidator)

	userUsecase := usecase.NewUserUsecase(mockRepo, mockHasher, mockTokenGen, MockValidator)

	user := &domain.User{
		Email:    "testAbc@gmail.com",
		Password: "passwordA@#$123",
	}

	mockHasher.On("EncryptPassword", mock.Anything).Return("passwordA@#$123", nil)
	mockRepo.On("GetUserByEmail", "testAbc@gmail.com").Return(nil, nil)
	mockRepo.On("AddUser", user).Return(user, nil)

	MockValidator.On("IsValidEmail", mock.Anything).Return(true)
	MockValidator.On("IsValidPassword", mock.Anything).Return(true)

	createdUser, err := userUsecase.CreateAdmin(user)

	assert.NoError(t, err)
	assert.Equal(t, "admin", createdUser.Role)

	mockRepo.AssertExpectations(t)

}

func TestCreateAdminRolePreAssignedAsUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockHasher := new(MockPasswordHasher)
	mockTokenGen := new(MockTokenGenerator)
	MockValidator := new(MockValidator)

	userUsecase := usecase.NewUserUsecase(mockRepo, mockHasher, mockTokenGen, MockValidator)

	user := &domain.User{
		Email:    "testAbc@gmail.com",
		Password: "passwordA@#$123",
		Role:     "user",
	}

	mockHasher.On("EncryptPassword", mock.Anything).Return("passwordA@#$123", nil)
	mockRepo.On("GetUserByEmail", "testAbc@gmail.com").Return(nil, nil)
	mockRepo.On("AddUser", user).Return(user, nil)

	MockValidator.On("IsValidEmail", mock.Anything).Return(true)
	MockValidator.On("IsValidPassword", mock.Anything).Return(true)

	createdUser, err := userUsecase.CreateAdmin(user)

	assert.NoError(t, err)
	assert.Equal(t, "admin", createdUser.Role, "expected created user role to be: %v, but got: %v", "admin", createdUser.Role)

	mockRepo.AssertExpectations(t)

}
