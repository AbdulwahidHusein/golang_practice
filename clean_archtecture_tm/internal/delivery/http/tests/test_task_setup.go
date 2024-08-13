package tests

import (
	"net/http/httptest"
	"task_managemet_api/cmd/task_manager/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	myHttp "task_managemet_api/cmd/task_manager/internal/delivery/http"
)

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
