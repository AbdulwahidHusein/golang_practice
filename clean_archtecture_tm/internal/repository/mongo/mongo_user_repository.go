package mongo

import (
	"context"
	"task_managemet_api/cmd/task_manager/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUserCollection interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
}

type MongoUserRepository struct {
	userCollection MongoUserCollection
}

func NewMongoUserRepository(userCollection MongoUserCollection) *MongoUserRepository {
	// userCollection := client.Database(database).Collection(collection)
	return &MongoUserRepository{userCollection}
}

func (r *MongoUserRepository) IsEmptyCollection(ctx context.Context) (bool, error) {
	count, err := r.userCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (r *MongoUserRepository) AddUser(user *domain.User) (*domain.User, error) {
	user.ID = primitive.NewObjectID()
	_, err := r.userCollection.InsertOne(context.TODO(), user)
	return user, err
}

func (r *MongoUserRepository) DeleteUser(id primitive.ObjectID) error {
	_, err := r.userCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func (r *MongoUserRepository) UpdateUser(id primitive.ObjectID, user *domain.User) (*domain.User, error) {
	_, err := r.userCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": user})
	if err != nil {
		return nil, err
	}
	var updatedUser domain.User
	err = r.userCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&updatedUser)
	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

func (r *MongoUserRepository) GetUSerById(id primitive.ObjectID) (*domain.User, error) {
	var user domain.User
	err := r.userCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *MongoUserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.userCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
