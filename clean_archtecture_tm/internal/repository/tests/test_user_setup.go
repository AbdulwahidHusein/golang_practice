package tests

import (
	"context"
	"task_managemet_api/cmd/task_manager/config"
	"task_managemet_api/cmd/task_manager/internal/repository"

	LocalMongo "task_managemet_api/cmd/task_manager/internal/repository/mongo"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
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
	clientOptions := options.Client().ApplyURI(config.GetEnvs()["MONGODB_URI"])
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
