package tests

import (
	"context"
	"task_managemet_api/cmd/task_manager/internal/domain"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//test the AddUser method

func (suite *MongoUserRepositoryTestSuite) TestAddUser_positive() {

	user := &domain.User{
		Email:    "test_user@gmail.com",
		Password: "test_password",
	}
	usr, err := suite.repo.AddUser(user)
	require.NoError(suite.T(), err, "Failed to add user")
	require.Equal(suite.T(), user.Email, usr.Email, "user email does not match")

	var result domain.User
	err = suite.userCollection.FindOne(context.TODO(), bson.M{"email": "test_user@gmail.com"}).Decode(&result)
	require.NoError(suite.T(), err, "Failed to find user in collection")

}

func (suite *MongoUserRepositoryTestSuite) TestAddUser_WithExistingId() {
	existingId := primitive.NewObjectID()

	user := &domain.User{
		ID:       existingId,
		Email:    "test_user@gmail.com",
		Password: "test_password",
	}
	_, err := suite.repo.AddUser(user)
	require.NoError(suite.T(), err, "should not return error adding user")

	addedUser, err := suite.repo.GetUSerById(existingId)

	suite.Equal(suite.T(), mongo.ErrNoDocuments, err)
	suite.Equal(suite.T(), nil, addedUser)
}

func (suite *MongoUserRepositoryTestSuite) TestAddUser_WithoutEmail() {
	user := &domain.User{
		Password: "test_password",
	}
	_, err := suite.repo.AddUser(user)
	require.Error(suite.T(), err, "Expected error adding user without email")

}

func (suite *MongoUserRepositoryTestSuite) TestAddUser_WithoutPassword() {
	user := &domain.User{
		Email: "test_@gmail.com",
	}
	_, err := suite.repo.AddUser(user)
	require.Error(suite.T(), err, "Expected error adding user without password")

}

// test the GetUser method

func (suite *MongoUserRepositoryTestSuite) TestGetUserByEmail_Dt_Cases() {

	user := &domain.User{
		Email:    "test_user@gmail.com",
		Password: "test_password",
	}
	_, err := suite.repo.AddUser(user)
	require.NoError(suite.T(), err, "Failed to add user")

	dbUser, err := suite.repo.GetUserByEmail("noemail@email.com")
	require.NoError(suite.T(), err, "Failed to get user")
	require.Equal(suite.T(), nil, dbUser, "user should be nil")

	dbUser, err = suite.repo.GetUserByEmail("test_user@gmail.com")
	require.NoError(suite.T(), err, "Failed to get user")
	require.Equal(suite.T(), user.Email, dbUser.Email, "user email does not match")
}

//Test getUSerById

func (suite *MongoUserRepositoryTestSuite) TestGetUserById_positive() {

	user := &domain.User{
		Email:    "test_user@gmail.com",
		Password: "test_password",
	}
	_, err := suite.repo.AddUser(user)
	require.NoError(suite.T(), err, "Failed to add user")

	dbUser, err := suite.repo.GetUSerById(user.ID)
	require.NoError(suite.T(), err, "Failed to get user")
	require.Equal(suite.T(), user.Email, dbUser.Email, "user email does not match")
}

// test UpdateUser

func (suite *MongoUserRepositoryTestSuite) TestUpdateUser_positive() {

	user := &domain.User{
		Email:    "test_user@gmail.com",
		Password: "test_password",
	}
	_, err := suite.repo.AddUser(user)
	require.NoError(suite.T(), err, "Failed to add user")

	tobeUpdated := &domain.User{
		ID:        user.ID,
		FirstName: "abdulwahid",
	}
	usr, err := suite.repo.UpdateUser(user.ID, tobeUpdated)
	require.NoError(suite.T(), err, "Failed to update user")
	require.Equal(suite.T(), "abdulwahid", usr.FirstName, "user first name does not match")
}

func (suite *MongoUserRepositoryTestSuite) TestUpdateUser_FieldsShouldNotChange() {

	user := &domain.User{
		Email:    "test_user@gmail.com",
		Password: "test_password",
	}
	_, err := suite.repo.AddUser(user)
	require.NoError(suite.T(), err, "Failed to add user")

	createdAt := user.CreatedAt
	email := user.Email
	password := user.Password

	tobeUpdated := &domain.User{
		ID:        user.ID,
		FirstName: "abdulwahid",
	}
	usr, err := suite.repo.UpdateUser(user.ID, tobeUpdated)
	require.NoError(suite.T(), err, "Failed to update user")
	require.Equal(suite.T(), createdAt, usr.CreatedAt, "user created at should not change")
	require.Equal(suite.T(), email, usr.Email, "user email should not change")
	require.Equal(suite.T(), password, usr.Password, "user password should not change")
	require.Equal(suite.T(), "abdulwahid", usr.FirstName, "user first name does not match")
}

// test DeleteUser

func (suite *MongoUserRepositoryTestSuite) TestDeleteUser() {

	user := &domain.User{
		Email:    "test_user@gmail.com",
		Password: "test_password",
	}
	_, err := suite.repo.AddUser(user)
	require.NoError(suite.T(), err, "Failed to add user")

	err = suite.repo.DeleteUser(user.ID)
	require.NoError(suite.T(), err, "Failed to delete user")
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(MongoUserRepositoryTestSuite))
}
