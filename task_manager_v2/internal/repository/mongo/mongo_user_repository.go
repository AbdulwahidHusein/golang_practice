package mongo

import (
	"context"
	"task_managemet_api/cmd/task_manager/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	userCollection *mongo.Collection
}

func NewMongoUserRepository(client *mongo.Client) *MongoUserRepository {
	userCollection := client.Database("task_manager_db").Collection("users")
	return &MongoUserRepository{userCollection}
}

func (r *MongoUserRepository) AddUser(user *domain.User) error {
	user.ID = primitive.NewObjectID()
	_, err := r.userCollection.InsertOne(context.TODO(), user)
	return err
}

func (r *MongoUserRepository) DeleteUser(id primitive.ObjectID) error {
	_, err := r.userCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func (r *MongoUserRepository) UpdateUser(id primitive.ObjectID, user *domain.User) error {
	_, err := r.userCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": user})
	return err
}

func (r *MongoUserRepository) GetUser(id primitive.ObjectID) (*domain.User, error) {
	var user domain.User
	err := r.userCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	return &user, err
}

func (r *MongoUserRepository) LoginUser(email string) (*domain.User, error) {
	var user domain.User
	err := r.userCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	return &user, err
}

func (r *MongoUserRepository) CheckUser(email string) (*domain.User, error) {
	var user domain.User
	err := r.userCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	return &user, err
}
