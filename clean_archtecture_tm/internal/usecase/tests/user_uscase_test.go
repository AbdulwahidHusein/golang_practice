package tests

import (
	"testing"

	"task_managemet_api/cmd/task_manager/internal/domain"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Test Cases
func (suite *UserUsecaseTestSuite) TestUserUsecase_AddUserPositive() {
	user := &domain.User{
		Email:    "testAbc@gmail.com",
		Password: "Stron@#$adgP@ssw0rd123",
	}

	// Mock the GetUserByEmail to return nil (indicating the user does not exist)
	suite.mockRepo.On("GetUserByEmail", mock.Anything).Return(nil, nil)

	// Mock the IsEmptyCollection to return false
	suite.mockRepo.On("IsEmptyCollection", mock.Anything).Return(false, nil)

	suite.mockHasher.On("EncryptPassword", mock.Anything).Return("Stron@#$adgP@ssw0rd123", nil)
	// Mock the AddUser to return the user
	suite.mockRepo.On("AddUser", user).Return(user, nil)
	suite.mockValidator.On("IsValidEmail", mock.Anything).Return(true)
	suite.mockValidator.On("IsValidPassword", mock.Anything).Return(true)

	createdUser, err := suite.userUsecase.AddUser(user)

	require.NoError(suite.T(), err)
	require.Equal(suite.T(), user.Email, createdUser.Email, "expected non-empty user email")
	require.Equal(suite.T(), user.Role, createdUser.Role, "expected non-empty user role")

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestUserUsecase_AddUserRoleAssignedAsAdmin() {
	user := &domain.User{
		Email:    "testAbc@gmail.com",
		Password: "Stron@#$adgP@ssw0rd123",
		Role:     "admin",
	}

	suite.mockRepo.On("GetUserByEmail", mock.Anything).Return(nil, nil)
	suite.mockRepo.On("IsEmptyCollection", mock.Anything).Return(false, nil)
	suite.mockHasher.On("EncryptPassword", mock.Anything).Return("Stron@#$adgP@ssw0rd123", nil)
	suite.mockRepo.On("AddUser", user).Return(user, nil)
	suite.mockValidator.On("IsValidEmail", mock.Anything).Return(true)
	suite.mockValidator.On("IsValidPassword", mock.Anything).Return(true)

	createdUser, err := suite.userUsecase.AddUser(user)

	require.NoError(suite.T(), err)
	require.Equal(suite.T(), user.Email, createdUser.Email, "expected non-empty user email")
	require.Equal(suite.T(), "user", createdUser.Role, "role must be user role but got %v", createdUser.Role)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestUserUsecase_AddUserUserAlreadyExist() {
	user := &domain.User{
		Email:    "testAbc@gmail.com",
		Password: "Stron@#$adgP@ssw0rd123",
	}

	suite.mockRepo.On("GetUserByEmail", mock.Anything).Return(user, nil)
	suite.mockValidator.On("IsValidEmail", mock.Anything).Return(true)
	suite.mockValidator.On("IsValidPassword", mock.Anything).Return(true)

	_, err := suite.userUsecase.AddUser(user)

	require.EqualError(suite.T(), err, "user with this email already exists")
}

func (suite *UserUsecaseTestSuite) TestFirstUserAdmin() {
	user := &domain.User{
		Email:    "testAbc@gmail.com",
		Password: "passwordA@#$123",
	}

	suite.mockRepo.On("GetUserByEmail", mock.Anything).Return(nil, nil)
	suite.mockRepo.On("IsEmptyCollection", mock.Anything).Return(true, nil)
	suite.mockRepo.On("AddUser", user).Return(user, nil)
	suite.mockHasher.On("EncryptPassword", mock.Anything).Return("passwordA@#$123", nil)
	suite.mockValidator.On("IsValidEmail", mock.Anything).Return(true)
	suite.mockValidator.On("IsValidPassword", mock.Anything).Return(true)

	createdUser, err := suite.userUsecase.AddUser(user)
	if err != nil {
		suite.T().Fatalf("expected no error while adding user, but got: %v", err)
	}
	require.Equal(suite.T(), "admin", createdUser.Role)

	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestUserUsecase_DeleteUserDtScenarios() {
	userID := primitive.NewObjectID()
	suite.mockHasher.On("ComparePassword", mock.Anything, mock.Anything).Return(nil)

	// Test unauthorized deletion
	err := suite.userUsecase.DeleteUser(userID, primitive.NewObjectID())
	require.EqualError(suite.T(), err, "unauthorized deletion", "expected unauthorized deletion error, but got: %v", err)

	// Test authorized deletion
	suite.mockRepo.On("DeleteUser", userID).Return(nil)
	err = suite.userUsecase.DeleteUser(userID, userID)
	require.NoError(suite.T(), err, "expected no error while deleting user, but got: %v", err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestUserUsecase_LoginUserPositive() {
	user := &domain.User{
		Email:    "test@example.com",
		Password: "$2a$12$ViW2yO/fbVtIbHDmPIgjNOEj6QqJqgWen33FFAhFT.0UCQhNDs1Ny", // Simulate a hashed password
		Role:     "user",
	}

	suite.mockTokenGen.On("CreateToken", mock.Anything, mock.Anything, mock.Anything).Return("access_token", "refresh_token", nil)
	suite.mockRepo.On("GetUserByEmail", mock.Anything).Return(user, nil)
	suite.mockHasher.On("ComparePassword", mock.Anything, mock.Anything).Return(nil)
	suite.mockValidator.On("IsValidEmail", mock.Anything).Return(true)
	suite.mockValidator.On("IsValidPassword", mock.Anything).Return(true)

	accessToken, refreshToken, err := suite.userUsecase.LoginUser(user.Email, user.Password)

	require.NoError(suite.T(), err)
	require.Equal(suite.T(), "access_token", accessToken)
	require.Equal(suite.T(), "refresh_token", refreshToken)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
