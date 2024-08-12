package tests

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	"github.com/mjarkk/mongomock"
// 	"github.com/stretchr/testify/require"
// 	"github.com/stretchr/testify/suite"
// 	"go.mongodb.org/mongo-driver/bson"

// 	"task_managemet_api/cmd/task_manager/internal/domain"
// 	"task_managemet_api/cmd/task_manager/internal/repository"
// 	LocalMongo "task_managemet_api/cmd/task_manager/internal/repository/mongo"
// )

// type MongoTaskRepositoryTestSuite struct {
// 	suite.Suite
// 	db             *mongomock.TestConnection
// 	repo           repository.TaskRepository
// 	taskCollection *mongomock.Collection
// }

// func (suite *MongoTaskRepositoryTestSuite) SetupSuite() {
// 	// Create a new mock database
// 	db := mongomock.NewDB()
// 	suite.taskCollection = db.Collection("tasks")

// 	// Initialize your repository with the mock collection
// 	suite.repo = LocalMongo.NewMongoTaskRepository(suite.taskCollection)
// }

// func (suite *MongoTaskRepositoryTestSuite) SetupTest() {
// 	err := suite.taskCollection.Drop(context.TODO())
// 	require.NoError(suite.T(), err, "Failed to drop test collection")
// }

// // TestAddTask tests the AddTask method
// func (suite *MongoTaskRepositoryTestSuite) TestAddTask() {
// 	task := &domain.Task{
// 		Title:       "Test Task",
// 		Description: "Test Description",
// 		DueDate:     time.Now(),
// 		Status:      "Pending",
// 	}

// 	err := suite.repo.AddTask(task)
// 	require.NoError(suite.T(), err, "Failed to add task")

// 	var result domain.Task
// 	err = suite.taskCollection.FindOne(context.TODO(), bson.M{"title": "Test Task"}).Decode(&result)
// 	require.NoError(suite.T(), err, "Failed to find task in collection")
// 	require.Equal(suite.T(), task.Title, result.Title, "Task title does not match")
// }

// // TestUpdateTask tests the UpdateTask method
// func (suite *MongoTaskRepositoryTestSuite) TestUpdateTask() {
// 	task := &domain.Task{
// 		Title:       "Test Task",
// 		Description: "Test Description",
// 		DueDate:     time.Now(),
// 		Status:      "Pending",
// 	}
// 	err := suite.repo.AddTask(task)
// 	require.NoError(suite.T(), err, "Failed to add task")

// 	err = suite.repo.UpdateTask(&domain.Task{
// 		ID:          task.ID,
// 		Title:       "Updated Task",
// 		Description: "Updated Description",
// 		DueDate:     time.Now(),
// 		Status:      "Pending",
// 	})
// 	require.NoError(suite.T(), err, "Failed to update task")
// 	var result domain.Task
// 	err = suite.taskCollection.FindOne(context.TODO(), bson.M{"_id": task.ID}).Decode(&result)
// 	require.NoError(suite.T(), err, "Failed to find task in collection")
// 	require.Equal(suite.T(), "Updated Task", result.Title, "Task title does not match")
// }

// // TestGetAllTasks tests the GetAllTasks method
// func (suite *MongoTaskRepositoryTestSuite) TestGetAllTasks() {
// 	task := &domain.Task{
// 		Title:       "Test Task update task",
// 		Description: "Test Description",
// 		DueDate:     time.Now(),
// 		Status:      "Pending",
// 	}
// 	err := suite.repo.AddTask(task)
// 	require.NoError(suite.T(), err, "Failed to add task")

// 	tasks, err := suite.repo.GetAllTasks()
// 	require.NoError(suite.T(), err, "Failed to get tasks")
// 	require.Equal(suite.T(), 1, len(tasks), "Expected 1 task, got %d", len(tasks))
// 	require.Equal(suite.T(), task.Title, tasks[0].Title, "Task title does not match")
// }

// // TestDeleteTask tests the DeleteTask method
// func (suite *MongoTaskRepositoryTestSuite) TestDeleteTask() {
// 	task := &domain.Task{
// 		Title:       "Test Task delete task",
// 		Description: "Test Description",
// 		DueDate:     time.Now(),
// 		Status:      "Pending",
// 	}
// 	err := suite.repo.AddTask(task)
// 	require.NoError(suite.T(), err, "Failed to add task")

// 	err = suite.repo.DeleteTask(task.ID.Hex())
// 	require.NoError(suite.T(), err, "Failed to delete task")

// 	tasks, err := suite.repo.GetAllTasks()
// 	require.NoError(suite.T(), err, "Failed to get tasks")
// 	require.Equal(suite.T(), 0, len(tasks), "Expected 0 tasks, got %d", len(tasks))
// }

// // Run the suite
// func TestMongoTaskRepositoryTestSuite(t *testing.T) {
// 	suite.Run(t, new(MongoTaskRepositoryTestSuite))
// }
