package tests

import (
	"context"
	"task_managemet_api/cmd/task_manager/internal/domain"
	"task_managemet_api/cmd/task_manager/internal/usecase"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mocks
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

// Test Suite
type UserUsecaseTestSuite struct {
	suite.Suite
	mockRepo      *MockUserRepository
	mockHasher    *MockPasswordHasher
	mockTokenGen  *MockTokenGenerator
	mockValidator *MockValidator
	userUsecase   usecase.UserUsecase
}

func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.mockRepo = new(MockUserRepository)
	suite.mockHasher = new(MockPasswordHasher)
	suite.mockTokenGen = new(MockTokenGenerator)
	suite.mockValidator = new(MockValidator)
	suite.userUsecase = usecase.NewUserUsecase(suite.mockRepo, suite.mockHasher, suite.mockTokenGen, suite.mockValidator)
}
