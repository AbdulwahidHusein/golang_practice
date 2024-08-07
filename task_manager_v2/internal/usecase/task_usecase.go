package usecase

import (
	"task_managemet_api/cmd/task_manager/internal/domain"

	"task_managemet_api/cmd/task_manager/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecase struct {
	taskRepository repository.TaskRepository
}

func NewTaskUseCase(taskRepository repository.TaskRepository) TaskUsecase {
	return TaskUsecase{
		taskRepository: taskRepository,
	}
}

func (u TaskUsecase) AddTask(task *domain.Task) error {
	return u.taskRepository.AddTask(task)
}

func (u TaskUsecase) GetTasks() ([]*domain.Task, error) {
	return u.taskRepository.GetAllTasks()
}

func (u TaskUsecase) GetTask(id string) (*domain.Task, error) {
	return u.taskRepository.GetTaskById(id)
}

func (u TaskUsecase) UpdateTask(taskId string, task *domain.Task) error {
	task.ID, _ = primitive.ObjectIDFromHex(taskId)
	return u.taskRepository.UpdateTask(task)
}

func (u TaskUsecase) DeleteTask(id string) error {
	return u.taskRepository.DeleteTask(id)
}

func (u TaskUsecase) GetDoneTasks() ([]*domain.Task, error) {
	return u.taskRepository.GetDoneTasks()
}

func (u TaskUsecase) GetUndoneTasks() ([]*domain.Task, error) {
	return u.taskRepository.GetUndoneTasks()
}
