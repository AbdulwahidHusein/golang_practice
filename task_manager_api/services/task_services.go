package services

import (
	"task_management_api/models"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	task_collection *mongo.Collection
}

func NewTaskService(client *mongo.Client) *TaskService {
	task_collection := client.Database("task_manager_db").Collection("tasks")
	return &TaskService{task_collection}
}

func (s *TaskService) AddTask(task *models.Task) error {
	task.ID = primitive.NewObjectID()
	_, err := s.task_collection.InsertOne(context.TODO(), task)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) GetTasks(userId string) ([]*models.Task, error) {
	var tasks []*models.Task
	cursor, err := s.task_collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (s *TaskService) GetTask(id string) (*models.Task, error) {
	var task models.Task
	oid, _ := primitive.ObjectIDFromHex(id)
	err := s.task_collection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) UpdateTask(task *models.Task) error {
	id := task.ID
	_, err := s.task_collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": task})
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) DeleteTask(id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	_, err := s.task_collection.DeleteOne(context.TODO(), bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) GetDoneTasks() ([]*models.Task, error) {
	var tasks []*models.Task
	cursor, err := s.task_collection.Find(context.TODO(), bson.M{"status": "done"})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (s *TaskService) GetUndoneTasks() ([]*models.Task, error) {
	var tasks []*models.Task
	cursor, err := s.task_collection.Find(context.TODO(), bson.M{"status": bson.M{"$ne": "done"}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}
