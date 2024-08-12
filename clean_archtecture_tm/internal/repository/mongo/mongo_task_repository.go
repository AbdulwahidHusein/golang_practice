package mongo

import (
	"context"
	"task_managemet_api/cmd/task_manager/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTaskRepository struct {
	taskCollection *mongo.Collection
}

func NewMongoTaskRepository(collection *mongo.Collection) *MongoTaskRepository {
	// taskCollection := client.Database(database).Collection(collection)
	return &MongoTaskRepository{collection}
}

func (r *MongoTaskRepository) AddTask(task *domain.Task) error {
	task.ID = primitive.NewObjectID()
	_, err := r.taskCollection.InsertOne(context.TODO(), task)
	return err
}

func (r *MongoTaskRepository) GetAllTasks() ([]*domain.Task, error) {
	var tasks []*domain.Task
	cursor, err := r.taskCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var task domain.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (r *MongoTaskRepository) GetTaskById(id string) (*domain.Task, error) {
	var task domain.Task
	oid, _ := primitive.ObjectIDFromHex(id)
	err := r.taskCollection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&task)
	return &task, err
}

func (r *MongoTaskRepository) UpdateTask(task *domain.Task) error {
	id := task.ID
	updateFields := bson.M{
		"title":       task.Title,
		"description": task.Description,
		"due_date":    task.DueDate,
		"status":      task.Status,
	}

	update := bson.M{"$set": updateFields}
	_, err := r.taskCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	return err
}

func (r *MongoTaskRepository) DeleteTask(id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	_, err := r.taskCollection.DeleteOne(context.TODO(), bson.M{"_id": oid})
	return err
}

func (r *MongoTaskRepository) GetTasksWithCriteria(criteria map[string]interface{}) ([]*domain.Task, error) {
	filter := bson.M{}
	for k, v := range criteria {
		filter[k] = v
	}

	var tasks []*domain.Task
	cursor, err := r.taskCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task *domain.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
