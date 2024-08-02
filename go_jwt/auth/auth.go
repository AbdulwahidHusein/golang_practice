package auth

import (
	"context"
	"net/http"

	"go_jwt/authUtils"
	"go_jwt/config"
	"go_jwt/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	UserCollection *mongo.Collection
}

func NewUserService(client *mongo.Client) *UserService {
	collection := client.Database("go_jwt").Collection("usersDb")

	// unique index on the email field
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic("Failed to create unique index: " + err.Error())
	}

	return &UserService{UserCollection: collection}
}

func (s *UserService) CreateUser(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if !authUtils.IsValidPassword(u.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character"})
		return
	}

	if !authUtils.IsValidEmail(u.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	enryptedPassword, err := authUtils.EncryptPassword(u.Password)
	u.Password = enryptedPassword
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
		return
	}

	u.ID = primitive.NewObjectID()
	_, err = s.UserCollection.InsertOne(context.TODO(), u)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, u)
}

func (s *UserService) GetUsers(c *gin.Context) {
	var users []*models.User
	cursor, err := s.UserCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users: " + err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode user data: " + err.Error()})
			return
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cursor error: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (s *UserService) LoginUser(c *gin.Context) {
	var guest models.User
	if err := c.BindJSON(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var realUser models.User
	err := s.UserCollection.FindOne(context.TODO(), bson.M{"email": guest.Email}).Decode(&realUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := authUtils.ComparePassword(realUser.Password, guest.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	tokenString, err := config.CreateToken(guest.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
