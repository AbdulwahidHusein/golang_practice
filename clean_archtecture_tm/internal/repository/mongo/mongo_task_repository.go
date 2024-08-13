package mongo

import (
	"context"
	"errors"
	"task_managemet_api/cmd/task_manager/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoTaskCollection interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}

type MongoTaskRepository struct {
	taskCollection MongoTaskCollection
}

func NewMongoTaskRepository(collection MongoTaskCollection) *MongoTaskRepository {
	// taskCollection := client.Database(database).Collection(collection)
	return &MongoTaskRepository{collection}
}

func (r *MongoTaskRepository) AddTask(task *domain.Task) error {
	task.ID = primitive.NewObjectID()
	if len(task.Description) > 1000 || len(task.Title) > 1000 {
		return errors.New("task title and description must be less than 1000 characters")
	}
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

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
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
	originalTask := domain.Task{}
	err := r.taskCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&originalTask)
	if err != nil {
		return err
	}
	update := bson.M{"$set": updateFields}
	_, err1 := r.taskCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	return err1
}

func (r *MongoTaskRepository) DeleteTask(id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	err := r.taskCollection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&domain.Task{})
	if err == mongo.ErrNoDocuments {
		return errors.New("task not found")
	}
	_, err1 := r.taskCollection.DeleteOne(context.TODO(), bson.M{"_id": oid})
	return err1
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
