package http

import (
	"encoding/json"
	"net/http"
	"task_managemet_api/cmd/task_manager/internal/domain"

	"github.com/gin-gonic/gin"
)

type TaskUsecase interface {
	AddTask(task *domain.Task) error
	GetTasks() ([]*domain.Task, error)
	GetTask(id string) (*domain.Task, error)
	UpdateTask(taskID string, task *domain.Task) error
	DeleteTask(id string) error
	GetTaskByStatus(status string) ([]*domain.Task, error)
}

type TaskHandler struct {
	TaskUsecase TaskUsecase
}

func NewTaskHandler(taskUsecase TaskUsecase) *TaskHandler {
	return &TaskHandler{
		TaskUsecase: taskUsecase,
	}
}

func (c *TaskHandler) CreateTask(ctx *gin.Context) {
	task := domain.Task{}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if task.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "title is required"})
		return
	}
	err := c.TaskUsecase.AddTask(&task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, task)
}

func (c *TaskHandler) GetTasks(ctx *gin.Context) {
	tasks, err := c.TaskUsecase.GetTasks()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskHandler) GetTask(ctx *gin.Context) {
	taskId := ctx.Param("id")
	task, err := c.TaskUsecase.GetTask(taskId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (c *TaskHandler) UpdateTask(ctx *gin.Context) {
	taskId := ctx.Param("id")
	task := domain.Task{}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.TaskUsecase.UpdateTask(taskId, &task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (c *TaskHandler) DeleteTask(ctx *gin.Context) {
	taskId := ctx.Param("id")

	err := c.TaskUsecase.DeleteTask(taskId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func (c *TaskHandler) GetTaskByStatus(ctx *gin.Context) {
	status := ctx.Param("status")
	tasks, err := c.TaskUsecase.GetTaskByStatus(status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}
