package repository

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
