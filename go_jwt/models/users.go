package models

import (
	"context"
	"net/http"

	"go_jwt/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string             `bson:"email" json:"email" gorm:"unique"`
	Password string             `bson:"password" json:"password" gorm:"not null"`
}

type UserService struct {
	UserCollection *mongo.Collection
}

func NewUserService(client *mongo.Client) *UserService {
	collection := client.Database("go_jwt").Collection("usersDb")

	// Ensure unique index on the email field
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}

	return &UserService{UserCollection: collection}
}

func (s *UserService) CreateUser(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.ID = primitive.NewObjectID()
	_, err := s.UserCollection.InsertOne(context.TODO(), u)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, u)
}

func (s *UserService) GetUsers(c *gin.Context) {
	var users []*User
	cursor, err := s.UserCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (s *UserService) LoginUser(c *gin.Context) {

	var user User

	if c.BindJSON(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := s.UserCollection.FindOne(context.TODO(), bson.M{"email": user.Email, "password": user.Password}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password", "email": user.Email, "password": user.Password})
		return
	}
	tokenString, err1 := config.CreateToken(user.Email)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err1.Error()})
		return
	}
	response := gin.H{
		"token": tokenString,
	}
	c.JSON(http.StatusOK, response)
}
