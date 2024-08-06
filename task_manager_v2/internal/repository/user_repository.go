package repository

import (
	"context"
	"task_managemet_api/cmd/task_manager/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServices struct {
	user_collection *mongo.Collection
}

func NewUserServices(client *mongo.Client) *UserServices {
	user_collection := client.Database("task_manager_db").Collection("users")
	return &UserServices{user_collection}
}

func (u *UserServices) AddUser(user *domain.User) error {
	user.ID = primitive.NewObjectID()
	_, err := u.user_collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServices) DeleteUser(id primitive.ObjectID) error {

	_, err := u.user_collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServices) UpdateUser(id primitive.ObjectID, user *domain.User) error {

	_, err := u.user_collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServices) GetUser(id primitive.ObjectID) (*domain.User, error) {

	var user domain.User
	err := u.user_collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserServices) LoginUser(email string) (*domain.User, error) {

	var user domain.User
	err := u.user_collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserServices) CheckUser(email string) (*domain.User, error) {

	var user domain.User
	err := u.user_collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
