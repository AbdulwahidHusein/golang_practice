package tests

import (
	"context"
	"strings"
	"task_managemet_api/cmd/task_manager/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TestAddTask tests the AddTask method
func (suite *MongoTaskRepositoryTestSuite) TestAddTask_Positive() {
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

func (suite *MongoTaskRepositoryTestSuite) TestAddTask_WithExistingId() {
	ExistingId := primitive.NewObjectID()
	task := &domain.Task{
		ID: ExistingId,
	}
	err := suite.repo.AddTask(task)
	require.NoError(suite.T(), err)

	// Check if the task was not added
	tasks, err := suite.repo.GetAllTasks()
	require.NoError(suite.T(), err)
	require.Len(suite.T(), tasks, 1)
	require.NotEqual(suite.T(), ExistingId, tasks[0].ID)
}

func (suite *MongoTaskRepositoryTestSuite) TestAddTask_LargeData() {
	description := strings.Repeat("a", 10*1024*1024) // 10 MB string

	task := &domain.Task{
		Title:       "Test Task",
		Description: string(description),
		DueDate:     time.Now(),
		Status:      "Pending",
	}

	err := suite.repo.AddTask(task)
	require.Error(suite.T(), err)

	tasks, err := suite.repo.GetAllTasks()
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), 0, len(tasks))
}

func (suite *MongoTaskRepositoryTestSuite) TestGetAllTasks_Positive() {
	task1 := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		DueDate:     time.Now(),
		Status:      "Pending",
	}
	task2 := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		DueDate:     time.Now(),
		Status:      "Pending",
	}
	require.NoError(suite.T(), suite.repo.AddTask(task1))
	require.NoError(suite.T(), suite.repo.AddTask(task2))

	tasks, err := suite.repo.GetAllTasks()
	require.NoError(suite.T(), err)
	require.Len(suite.T(), tasks, 2)
}

func (suite *MongoTaskRepositoryTestSuite) TestGetAllTasks_EmptyCollection() {
	tasks, err := suite.repo.GetAllTasks()
	require.NoError(suite.T(), err)
	require.Empty(suite.T(), tasks)
}

func (suite *MongoTaskRepositoryTestSuite) TestGetTaskById_Positive() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		DueDate:     time.Now(),
		Status:      "Pending",
	}
	require.NoError(suite.T(), suite.repo.AddTask(task))

	addedTask, err := suite.repo.GetTaskById(task.ID.Hex())
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), task.Title, addedTask.Title)
}

func (suite *MongoTaskRepositoryTestSuite) TestGetTaskById_TaskNotFound() {
	taskId := primitive.NewObjectID()
	task, err := suite.repo.GetTaskById(taskId.Hex())
	require.NoError(suite.T(), err)
	require.Nil(suite.T(), task)
}

func (suite *MongoTaskRepositoryTestSuite) TestUpdateTask_Positive() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		DueDate:     time.Now(),
		Status:      "Pending",
	}
	require.NoError(suite.T(), suite.repo.AddTask(task))
	prevId := task.ID

	newId := primitive.NewObjectID()
	task.Description = "Updated Description"
	task.Title = "Updated Title"
	// task.ID = newId

	err := suite.repo.UpdateTask(task)
	require.NoError(suite.T(), err)
	require.NotEqual(suite.T(), newId, task.ID)
	require.Equal(suite.T(), prevId, task.ID)
	require.Equal(suite.T(), "Updated Title", task.Title)
}

func (suite *MongoTaskRepositoryTestSuite) TestUpdateTask_ID_Unassinable() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		DueDate:     time.Now(),
		Status:      "Pending",
	}
	require.NoError(suite.T(), suite.repo.AddTask(task))

	newId := primitive.NewObjectID()
	task.Description = "Updated Description"
	task.Title = "Updated Title"
	task.ID = newId

	err := suite.repo.UpdateTask(task)
	require.Error(suite.T(), err)
}

func (suite *MongoTaskRepositoryTestSuite) TestDeleteTask_Positive() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		DueDate:     time.Now(),
		Status:      "Pending",
	}
	require.NoError(suite.T(), suite.repo.AddTask(task))

	err := suite.repo.DeleteTask(task.ID.Hex())
	require.NoError(suite.T(), err)

	tasks, err := suite.repo.GetAllTasks()
	require.NoError(suite.T(), err)
	require.Len(suite.T(), tasks, 0)
}

func (suite *MongoTaskRepositoryTestSuite) TestDeleteTask_TaskNotFound() {
	err := suite.repo.DeleteTask("123")
	require.Error(suite.T(), err)

}

func (suite *MongoTaskRepositoryTestSuite) TestGetTasksWithCriteria_positive() {

	task1 := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		DueDate:     time.Now(),
		Status:      "Pending",
	}
	task2 := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		DueDate:     time.Now(),
		Status:      "done",
	}
	require.NoError(suite.T(), suite.repo.AddTask(task1))
	require.NoError(suite.T(), suite.repo.AddTask(task2))

	criteria := map[string]interface{}{
		"title":  "Test Task",
		"status": "Pending",
	}
	tasks, err := suite.repo.GetTasksWithCriteria(criteria)
	require.NoError(suite.T(), err)
	require.Len(suite.T(), tasks, 1)
}

func (suite *MongoTaskRepositoryTestSuite) TestGetTasksWithCriteria_InvalidFieldInCriteria() {
	task := &domain.Task{
		Title: "Test Task",
	}
	err := suite.repo.AddTask(task)
	require.NoError(suite.T(), err)

	criteria := map[string]interface{}{
		"nonExistentField": "someValue",
	}

	tasks, err := suite.repo.GetTasksWithCriteria(criteria)
	require.NoError(suite.T(), err)
	require.Empty(suite.T(), tasks)
}

// Run the suite
func TestMongoTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(MongoTaskRepositoryTestSuite))
}
