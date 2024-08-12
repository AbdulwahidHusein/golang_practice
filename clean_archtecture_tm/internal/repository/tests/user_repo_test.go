package tests

import (
	"context"

	"task_managemet_api/cmd/task_manager/config"
	"task_managemet_api/cmd/task_manager/internal/domain"
	LocalMongo "task_managemet_api/cmd/task_manager/internal/repository/mongo"

	"task_managemet_api/cmd/task_manager/internal/repository"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoUserRepositoryTestSuite struct {
	suite.Suite
	client         *mongo.Client
	repo           repository.UserRepository // Use the interface type here
	userCollection *mongo.Collection
}

// SetupSuite runs once before all tests
func (suite *MongoUserRepositoryTestSuite) SetupSuite() {
	clientOptions := options.Client().ApplyURI(config.GetEnvs()["MONGO_URI"])
	client, err := mongo.Connect(context.TODO(), clientOptions)
	require.NoError(suite.T(), err, "Failed to connect to MongoDB")

	err = client.Ping(context.TODO(), readpref.Primary())
	require.NoError(suite.T(), err, "Failed to ping MongoDB")

	suite.client = client
	suite.userCollection = client.Database("task_manager_db_test").Collection("users")
	suite.repo = LocalMongo.NewMongoUserRepository(suite.userCollection)
}

// TearDownSuite runs once after all tests
func (suite *MongoUserRepositoryTestSuite) TearDownSuite() {
	err := suite.client.Disconnect(context.TODO())
	require.NoError(suite.T(), err, "Failed to disconnect from MongoDB")
}

// SetupTest runs before each test to ensure the collection is clean
func (suite *MongoUserRepositoryTestSuite) SetupTest() {
	err := suite.userCollection.Drop(context.TODO())
	require.NoError(suite.T(), err, "Failed to drop test collection")
}

//test the AddUser method

func (suite *MongoUserRepositoryTestSuite) TestAddUser() {

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

// test the GetUser method

func (suite *MongoUserRepositoryTestSuite) TestGetUserByEmail() {

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

func (suite *MongoUserRepositoryTestSuite) TestGetUserById() {

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

func (suite *MongoUserRepositoryTestSuite) TestUpdateUser() {

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
