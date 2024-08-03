package controllers

import (
	"encoding/json"
	"net/http"
	"task_management_api/models"
	"task_management_api/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskDbOperator interface {
	AddTask(task *models.Task) error
	GetTasks(userId primitive.ObjectID) ([]*models.Task, error)
	GetTask(id string) (*models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id string) error
	GetDoneTasks(primitive.ObjectID) ([]*models.Task, error)
	GetUndoneTasks(userId primitive.ObjectID) ([]*models.Task, error)
}

type TaskController struct {
	taskService TaskDbOperator
}

func NewTaskController(service TaskDbOperator) *TaskController {
	return &TaskController{service}
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	task := models.Task{}
	userId, err1 := utils.GetUSerIdFormToken(ctx)
	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	task.UserId = userId
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
	userId, err := utils.GetUSerIdFormToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tasks, err := c.taskService.GetTasks(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskController) GetTask(ctx *gin.Context) {
	userId, err := utils.GetUSerIdFormToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	taskId := ctx.Param("id")
	task, err := c.taskService.GetTask(taskId)

	if task.UserId != userId {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	taskId := ctx.Param("id")

	taskWithId, err := c.taskService.GetTask(taskId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "task not found"})
		return
	}
	ModifiedTask := models.Task{}
	userId, err1 := utils.GetUSerIdFormToken(ctx)

	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	if err := json.NewDecoder(ctx.Request.Body).Decode(&ModifiedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	PremitiveTaskID, err2 := primitive.ObjectIDFromHex(taskId)

	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}
	ModifiedTask.ID = PremitiveTaskID

	if taskWithId.UserId != userId {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "You dont have permission to update this task"})
		return
	}
	err3 := c.taskService.UpdateTask(&ModifiedTask)

	if err3 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err3.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ModifiedTask)
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	taskId := ctx.Param("id")
	userId, err1 := utils.GetUSerIdFormToken(ctx)
	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	task, err2 := c.taskService.GetTask(taskId)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return
	}
	if task.UserId != userId {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	err := c.taskService.DeleteTask(taskId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func (c *TaskController) GetDoneTasks(ctx *gin.Context) {
	userID, err := utils.GetUSerIdFormToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tasks, err := c.taskService.GetDoneTasks(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskController) GetUndoneTasks(ctx *gin.Context) {
	userId, err := utils.GetUSerIdFormToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tasks, err := c.taskService.GetUndoneTasks(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}
