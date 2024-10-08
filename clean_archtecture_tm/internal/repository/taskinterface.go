package repository

import (
	"task_managemet_api/cmd/task_manager/internal/domain"
)

type TaskRepository interface {
	AddTask(task *domain.Task) error
	GetAllTasks() ([]*domain.Task, error)
	GetTaskById(id string) (*domain.Task, error)
	UpdateTask(task *domain.Task) error
	DeleteTask(id string) error
	GetTasksWithCriteria(criteria map[string]interface{}) ([]*domain.Task, error)
}
