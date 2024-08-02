package controllers

import (
	"encoding/json"
	"net/http"
	"task_management_api/models"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskDbOperator interface {
	AddTask(task *models.Task) error
	GetTasks() ([]*models.Task, error)
	GetTask(id string) (*models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id string) error
	GetDoneTasks() ([]*models.Task, error)
	GetUndoneTasks() ([]*models.Task, error)
}

type TaskController struct {
	taskService TaskDbOperator
}

func NewTaskController(service TaskDbOperator) *TaskController {
	return &TaskController{service}
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	task := models.Task{}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.taskService.AddTask(&task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, task)
}

func (c *TaskController) GetTasks(ctx *gin.Context) {
	tasks, err := c.taskService.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskController) GetTask(ctx *gin.Context) {
	taskId := ctx.Param("id")
	task, err := c.taskService.GetTask(taskId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	task := models.Task{}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := c.taskService.UpdateTask(&task)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	taskId := ctx.Param("id")
	err := c.taskService.DeleteTask(taskId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func (c *TaskController) GetDoneTasks(ctx *gin.Context) {
	tasks, err := c.taskService.GetDoneTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskController) GetUndoneTasks(ctx *gin.Context) {
	tasks, err := c.taskService.GetUndoneTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}
