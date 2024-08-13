package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"task_managemet_api/cmd/task_manager/internal/domain"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func (suite *TaskHandlerSuite) TestCreateTask_Positive() {
	// should be succesful
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

func (suite *TaskHandlerSuite) TestCreateTask_Empty() {

	// Calling the testing server without data
	response, err := http.Post(fmt.Sprintf("%s/tasks", suite.testingServer.URL), "application/json", bytes.NewBuffer([]byte{}))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusBadRequest, response.StatusCode, "expected status code %d got %d", http.StatusBadRequest, response.StatusCode)
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *TaskHandlerSuite) TestCreateTask_NoTitle() {
	task := &domain.Task{
		Description: "This is a test task",
	}
	// suite.taskUsecase.On("AddTask", task).Return(nil)

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

	suite.Equal(http.StatusBadRequest, response.StatusCode, "expected status code %d got %d", http.StatusCreated, response.StatusCode)
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

func (suite *TaskHandlerSuite) TestGetTask_NoId() {

	// Calling the testing server
	response, err := http.Get(fmt.Sprintf("%s/tasks/%s", suite.testingServer.URL, ""))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusInternalServerError, response.StatusCode, "expected status code %d got %d", http.StatusBadRequest, response.StatusCode)
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

func (suite *TaskHandlerSuite) TestUpdateTask_EmptyTask() {

	// Marshalling

	requestBody, err := json.Marshal(&domain.Task{})
	suite.NoError(err, "can not marshal struct to json")

	suite.taskUsecase.On("UpdateTask", mock.Anything, mock.Anything).Return(errors.New("some error"))

	// Calling the testing server
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/tasks/%s", suite.testingServer.URL, "abc"), bytes.NewBuffer(requestBody))
	suite.NoError(err, "can not create request")
	req.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusInternalServerError, response.StatusCode, "expected status code %d got %d", http.StatusBadRequest, response.StatusCode)

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
