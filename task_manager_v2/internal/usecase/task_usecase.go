package usecase

import (
	"task_managemet_api/cmd/task_manager/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepository interface {
	AddTask(task *domain.Task) error
	GetTasks(userId primitive.ObjectID) ([]*domain.Task, error)
	GetTask(id string) (*domain.Task, error)
	UpdateTask(task *domain.Task) error
	DeleteTask(id string) error
	GetDoneTasks(userId primitive.ObjectID) ([]*domain.Task, error)
	GetUndoneTasks(userId primitive.ObjectID) ([]*domain.Task, error)
}

type TaskUsecase struct {
	taskRepository TaskRepository
}

func NewTaskUseCase(taskRepository TaskRepository) TaskUsecase {
	return TaskUsecase{
		taskRepository: taskRepository,
	}
}

func (u TaskUsecase) AddTask(task *domain.Task) error {
	return u.taskRepository.AddTask(task)
}

func (u TaskUsecase) GetTasks(userId primitive.ObjectID) ([]*domain.Task, error) {
	return u.taskRepository.GetTasks(userId)
}

func (u TaskUsecase) GetTask(id string) (*domain.Task, error) {
	return u.taskRepository.GetTask(id)
}

func (u TaskUsecase) UpdateTask(task *domain.Task) error {
	return u.taskRepository.UpdateTask(task)
}

func (u TaskUsecase) DeleteTask(id string) error {
	return u.taskRepository.DeleteTask(id)
}

func (u TaskUsecase) GetDoneTasks(userId primitive.ObjectID) ([]*domain.Task, error) {
	return u.taskRepository.GetDoneTasks(userId)
}

func (u TaskUsecase) GetUndoneTasks(userId primitive.ObjectID) ([]*domain.Task, error) {
	return u.taskRepository.GetUndoneTasks(userId)
}
