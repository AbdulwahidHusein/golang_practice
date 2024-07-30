package controllers

import (
	"encoding/json"
	"net/http"
	"task_management_api/models"
	"task_management_api/services"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	taskService *services.TaskService
}

func NewTaskController(service *services.TaskService) *TaskController {
	return &TaskController{taskService: service}
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	task := models.Task{}

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err := c.taskService.AddTask(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := c.taskService.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func (c *TaskController) GetTask(w http.ResponseWriter, r *http.Request) {
	taskId := r.URL.Query().Get("id")
	task, err := c.taskService.GetTask(taskId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func (c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	task := models.Task{}

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := c.taskService.UpdateTask(task.ID, &task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}
