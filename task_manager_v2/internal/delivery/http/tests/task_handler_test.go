package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	myHttp "task_managemet_api/cmd/task_manager/internal/delivery/http"
	"task_managemet_api/cmd/task_manager/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Mock for TaskRepository
type MockTaskUsecase struct {
	mock.Mock
}

func (m *MockTaskUsecase) AddTask(task *domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskUsecase) GetTasks() ([]*domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Task), args.Error(1)
}

func (m *MockTaskUsecase) GetTask(id string) (*domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTaskUsecase) UpdateTask(taskID string, task *domain.Task) error {
	args := m.Called(taskID, task)
	return args.Error(0)
}

func (m *MockTaskUsecase) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskUsecase) GetTaskByStatus(status string) ([]*domain.Task, error) {
	args := m.Called(status)
	return args.Get(0).([]*domain.Task), args.Error(1)
}

type TaskHandlerSuite struct {
	suite.Suite
	taskUsecase   *MockTaskUsecase
	handler       *myHttp.TaskHandler
	testingServer *httptest.Server
}

func (suite *TaskHandlerSuite) SetupTest() {
	suite.taskUsecase = new(MockTaskUsecase)
	suite.handler = myHttp.NewTaskHandler(suite.taskUsecase)

	router := gin.Default()
	r := router.Group("/tasks")
	r.POST("/", suite.handler.CreateTask)
	r.GET("/", suite.handler.GetTasks)
	r.GET("/:id", suite.handler.GetTask)
	r.PUT("/:id", suite.handler.UpdateTask)
	r.DELETE("/:id", suite.handler.DeleteTask)
	r.GET("/status/:status", suite.handler.GetTaskByStatus)

	suite.testingServer = httptest.NewServer(router)
}

func (suite *TaskHandlerSuite) TearDownSuite() {
	defer suite.testingServer.Close()
}

func (suite *TaskHandlerSuite) TestCreateTask_Positive() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "This is a test task",
	}
	suite.taskUsecase.On("AddTask", task).Return(nil)

	// Marshalling
	requestBody, err := json.Marshal(task)
	suite.NoError(err, "can not marshal struct to json")

	// Calling the testing server
	response, err := http.Post(fmt.Sprintf("%s/tasks", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	var responseBody domain.Task
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err, "can not unmarshal response body")

	suite.Equal(http.StatusCreated, response.StatusCode, "expected status code %d got %d", http.StatusCreated, response.StatusCode)
	suite.Equal(task.Title, responseBody.Title, "expected title %s got %s", task.Title, responseBody.Title)
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *TaskHandlerSuite) TestCreateTask_UsecaseReturnError() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "This is a test task",
	}
	suite.taskUsecase.On("AddTask", task).Return(errors.New("some error"))

	// Marshalling
	requestBody, err := json.Marshal(task)
	suite.NoError(err, "can not marshal struct to json")

	// Calling the testing server
	response, err := http.Post(fmt.Sprintf("%s/tasks", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusInternalServerError, response.StatusCode, "expected status code %d got %d", http.StatusInternalServerError, response.StatusCode)
}

func (suite *TaskHandlerSuite) TestGetTasks_Positive() {
	tasks := []*domain.Task{
		{Title: "Task 1", Description: "Description 1"},
		{Title: "Task 2", Description: "Description 2"},
	}
	suite.taskUsecase.On("GetTasks").Return(tasks, nil)

	// Calling the testing server
	response, err := http.Get(fmt.Sprintf("%s/tasks", suite.testingServer.URL))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	var responseBody []*domain.Task
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err, "can not unmarshal response body")

	suite.Equal(http.StatusOK, response.StatusCode, "expected status code %d got %d", http.StatusOK, response.StatusCode)
	suite.Len(responseBody, len(tasks), "expected %d tasks, got %d", len(tasks), len(responseBody))
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *TaskHandlerSuite) TestGetTask_Positive() {
	task := &domain.Task{
		Title:       "Test Task",
		Description: "This is a test task",
	}
	taskID := "some-task-id"
	suite.taskUsecase.On("GetTask", taskID).Return(task, nil)

	// Calling the testing server
	response, err := http.Get(fmt.Sprintf("%s/tasks/%s", suite.testingServer.URL, taskID))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	var responseBody domain.Task
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err, "can not unmarshal response body")

	suite.Equal(http.StatusOK, response.StatusCode, "expected status code %d got %d", http.StatusOK, response.StatusCode)
	suite.Equal(task.Title, responseBody.Title, "expected title %s got %s", task.Title, responseBody.Title)
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *TaskHandlerSuite) TestUpdateTask_Positive() {
	task := &domain.Task{
		Title:       "Updated Task",
		Description: "This task has been updated",
	}
	taskID := "some-task-id"
	suite.taskUsecase.On("UpdateTask", taskID, task).Return(nil)

	// Marshalling
	requestBody, err := json.Marshal(task)
	suite.NoError(err, "can not marshal struct to json")

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/tasks/%s", suite.testingServer.URL, taskID), bytes.NewBuffer(requestBody))
	suite.NoError(err, "can not create request")
	req.Header.Set("Content-Type", "application/json")

	// Calling the testing server
	response, err := http.DefaultClient.Do(req)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusOK, response.StatusCode, "expected status code %d got %d", http.StatusOK, response.StatusCode)
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *TaskHandlerSuite) TestDeleteTask_Positive() {
	taskID := "some-task-id"
	suite.taskUsecase.On("DeleteTask", taskID).Return(nil)

	// Calling the testing server
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/tasks/%s", suite.testingServer.URL, taskID), nil)
	suite.NoError(err, "can not create request")

	response, err := http.DefaultClient.Do(req)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusOK, response.StatusCode, "expected status code %d got %d", http.StatusOK, response.StatusCode)
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *TaskHandlerSuite) TestGetTaskByStatus_Positive() {
	status := "completed"
	tasks := []*domain.Task{
		{Title: "Completed Task 1", Description: "Description 1"},
		{Title: "Completed Task 2", Description: "Description 2"},
	}
	suite.taskUsecase.On("GetTaskByStatus", status).Return(tasks, nil)

	// Calling the testing server
	response, err := http.Get(fmt.Sprintf("%s/tasks/status/%s", suite.testingServer.URL, status))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	var responseBody []*domain.Task
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err, "can not unmarshal response body")

	suite.Equal(http.StatusOK, response.StatusCode, "expected status code %d got %d", http.StatusOK, response.StatusCode)
	suite.Len(responseBody, len(tasks), "expected %d tasks, got %d", len(tasks), len(responseBody))
	suite.taskUsecase.AssertExpectations(suite.T())
}

func TestTaskHandlerSuite(t *testing.T) {
	suite.Run(t, new(TaskHandlerSuite))
}
