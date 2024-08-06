package http

import (
	"encoding/json"
	"net/http"
	"task_managemet_api/cmd/task_manager/internal/domain"
	"task_managemet_api/cmd/task_manager/pkg/security"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecase interface {
	AddTask(task *domain.Task) error
	GetTasks(userId primitive.ObjectID) ([]*domain.Task, error)
	GetTask(id string) (*domain.Task, error)
	UpdateTask(task *domain.Task) error
	DeleteTask(id string) error
	GetDoneTasks(userId primitive.ObjectID) ([]*domain.Task, error)
	GetUndoneTasks(userId primitive.ObjectID) ([]*domain.Task, error)
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
	userId, err1 := security.GetUSerIdFormToken(ctx)
	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	task.UserId = userId
	if err := json.NewDecoder(ctx.Request.Body).Decode(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	userId, err := security.GetUSerIdFormToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tasks, err := c.TaskUsecase.GetTasks(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskHandler) GetTask(ctx *gin.Context) {
	userId, err := security.GetUSerIdFormToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	taskId := ctx.Param("id")
	task, err := c.TaskUsecase.GetTask(taskId)

	if task.UserId != userId {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (c *TaskHandler) UpdateTask(ctx *gin.Context) {
	taskId := ctx.Param("id")

	taskWithId, err := c.TaskUsecase.GetTask(taskId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "task not found"})
		return
	}
	ModifiedTask := domain.Task{}
	userId, err1 := security.GetUSerIdFormToken(ctx)

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
	err3 := c.TaskUsecase.UpdateTask(&ModifiedTask)

	if err3 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err3.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ModifiedTask)
}

func (c *TaskHandler) DeleteTask(ctx *gin.Context) {
	taskId := ctx.Param("id")
	userId, err1 := security.GetUSerIdFormToken(ctx)
	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	task, err2 := c.TaskUsecase.GetTask(taskId)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return
	}
	if task.UserId != userId {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	err := c.TaskUsecase.DeleteTask(taskId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func (c *TaskHandler) GetDoneTasks(ctx *gin.Context) {
	userID, err := security.GetUSerIdFormToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tasks, err := c.TaskUsecase.GetDoneTasks(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskHandler) GetUndoneTasks(ctx *gin.Context) {
	userId, err := security.GetUSerIdFormToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tasks, err := c.TaskUsecase.GetUndoneTasks(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}
