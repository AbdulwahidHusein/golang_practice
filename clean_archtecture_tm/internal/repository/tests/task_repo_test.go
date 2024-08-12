package tests

import (
	"context"
	"fmt"
	"os"
	"task_managemet_api/cmd/task_manager/config"
	"task_managemet_api/cmd/task_manager/internal/domain"
	"task_managemet_api/cmd/task_manager/internal/repository"

	MongoRepo "task_managemet_api/cmd/task_manager/internal/repository/mongo"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoTaskRepositoryTestSuite struct {
	suite.Suite
	client         *mongo.Client
	repo           repository.TaskRepository // Use the interface type here
	taskCollection *mongo.Collection
}

// SetupSuite runs once before all tests
func (suite *MongoTaskRepositoryTestSuite) SetupSuite() {
	cwd, _ := os.Getwd()
	fmt.Println("test repo file runs in a ", cwd)
	clientOptions := options.Client().ApplyURI(config.GetEnvs()["MONGODB_URI"])
	client, err := mongo.Connect(context.TODO(), clientOptions)
	require.NoError(suite.T(), err)

	err = client.Ping(context.TODO(), readpref.Primary())
	require.NoError(suite.T(), err)

	suite.client = client

	suite.taskCollection = suite.client.Database("task_manager_db").Collection("tasks")
	suite.repo = MongoRepo.NewMongoTaskRepository(suite.taskCollection) // Instantiate correctly
}

// TearDownSuite runs once after all tests
func (suite *MongoTaskRepositoryTestSuite) TearDownSuite() {
	err := suite.client.Disconnect(context.TODO())
	require.NoError(suite.T(), err)
}

// SetupTest runs before each test to ensure the collection is clean
func (suite *MongoTaskRepositoryTestSuite) SetupTest() {
	err := suite.taskCollection.Drop(context.TODO())
	require.NoError(suite.T(), err)
}

// TestAddTask tests the AddTask method
func (suite *MongoTaskRepositoryTestSuite) TestAddTask() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		DueDate:     time.Now(),
		Status:      "Pending",
	}

	err := suite.repo.AddTask(task)
	require.NoError(suite.T(), err)

	var result domain.Task
	err = suite.taskCollection.FindOne(context.TODO(), bson.M{"title": "Test Task"}).Decode(&result)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), task.Title, result.Title)
}

// Run the suite
func TestMongoTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(MongoTaskRepositoryTestSuite))
}
