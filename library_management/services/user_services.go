package services

import (
	"context"
	"library_management/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UerService struct {
	user_collection *mongo.Collection
}

func NewUserService(client *mongo.Client) *UerService {
	user_collection := client.Database("library_management").Collection("users")
	return &UerService{user_collection}
}

func (u *UerService) AddUser(user *models.Member) error {
	_, err := u.user_collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UerService) GetUser(id primitive.ObjectID) (*models.Member, error) {
	var user models.Member
	err := u.user_collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UerService) GetUsers() ([]*models.Member, error) {
	var users []*models.Member
	cursor, err := u.user_collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var user models.Member
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u *UerService) UpdateUser(id primitive.ObjectID, user *models.Member) error {
	_, err := u.user_collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": user})
	if err != nil {
		return err
	}
	return nil
}

func (u *UerService) DeleteUser(id primitive.ObjectID) error {
	_, err := u.user_collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
