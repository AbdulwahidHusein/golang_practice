package tests

import (
	"testing"
	"time"

	"task_managemet_api/cmd/task_manager/internal/domain"
	"task_managemet_api/cmd/task_manager/internal/usecase"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Mock for UserRepository
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) AddTask(task *domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) GetAllTasks() ([]*domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetTaskById(id string) (*domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(task *domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskRepository) GetTasksWithCriteria(criteria map[string]interface{}) ([]*domain.Task, error) {
	args := m.Called(criteria)
	return args.Get(0).([]*domain.Task), args.Error(1)
}

func TestTaskUsecase(t *testing.T) {
	mockTaskRepository := new(MockTaskRepository)
	mockTaskRepository.On("AddTask", mock.Anything).Return(nil)
	mockTaskRepository.On("GetAllTasks").Return(&domain.Task{}, nil)
	mockTaskRepository.On("GetTaskById", mock.Anything).Return([]*domain.Task{}, nil)
	mockTaskRepository.On("UpdateTask", mock.Anything).Return(nil)
	mockTaskRepository.On("DeleteTask", mock.Anything).Return(nil)
	mockTaskRepository.On("GetTasksWithCriteria", mock.Anything).Return([]*domain.Task{}, nil)

	taskUsecase := usecase.NewTaskUseCase(mockTaskRepository)

	task := &domain.Task{
		Title:       "Test Task",
		Description: "Test Description",
		DueDate:     time.Now(),
	}
	err := taskUsecase.AddTask(task)
	require.NoError(t, err, "Failed to add task")

}

func TestAddTak_positive(t *testing.T) {
	mockTaskRepository := new(MockTaskRepository)
	mockTaskRepository.On("AddTask", mock.Anything).Return(nil)
	require.NoError(t, mockTaskRepository.AddTask(&domain.Task{}))
}

func TestGetAllTasks(t *testing.T) {
	mockTaskRepository := new(MockTaskRepository)
	mockTaskRepository.On("GetAllTasks").Return([]*domain.Task{}, nil)
	_, err := mockTaskRepository.GetAllTasks()
	require.NoError(t, err)
}

func TestGetTaskById(t *testing.T) {
	mockTaskRepository := new(MockTaskRepository)
	mockTaskRepository.On("GetTaskById", mock.Anything).Return(&domain.Task{}, nil)
	_, err := mockTaskRepository.GetTaskById("testid")
	require.NoError(t, err)
}

func TestGetTasksWithCriteria(t *testing.T) {
	mockTaskRepository := new(MockTaskRepository)
	mockTaskRepository.On("GetTasksWithCriteria", mock.Anything).Return([]*domain.Task{}, nil)
	_, err := mockTaskRepository.GetTasksWithCriteria(map[string]interface{}{})
	require.NoError(t, err)
}
