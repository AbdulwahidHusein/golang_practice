package controllers

import (
	"net/http"
	"task_management_api/models"

	"task_management_api/auth"

	"task_management_api/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDBOperator interface {
	AddUser(user *models.User) error
	DeleteUser(id primitive.ObjectID) error
	UpdateUser(id primitive.ObjectID, user *models.User) error
	GetUser(id primitive.ObjectID) (*models.User, error)
	LoginUser(email string) (*models.User, error)
	CheckUser(email string) (*models.User, error)
}

type UserController struct {
	UserDBOperator UserDBOperator
}

func NewUSerController(userDBOperator UserDBOperator) *UserController {
	return &UserController{
		UserDBOperator: userDBOperator,
	}
}

func (u *UserController) AddUser(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)
	email := user.Email
	password := user.Password

	if !auth.IsValidEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}
	if !auth.IsValidPassword(password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character"})
		return
	}
	if user, _ := u.UserDBOperator.CheckUser(email); user != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}
	hashedPassword, _ := auth.EncryptPassword(password)
	user.Password = hashedPassword

	err := u.UserDBOperator.AddUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (u *UserController) LoginUser(c *gin.Context) {
	var guest models.User
	c.ShouldBindJSON(&guest)
	dbUser, err := u.UserDBOperator.LoginUser(guest.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	hashedPassword := dbUser.Password
	if err := auth.ComparePassword(hashedPassword, guest.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}
	tokenString, err := auth.CreateToken(dbUser.ID.Hex(), dbUser.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (u *UserController) UpdateUser(c *gin.Context) {
	var user models.User
	visitorId, err1 := utils.GetUSerIdFormToken(c)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	c.ShouldBindJSON(&user)
	user.ID = visitorId

	err := u.UserDBOperator.UpdateUser(user.ID, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	visitorId, err := utils.GetUSerIdFormToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, err := primitive.ObjectIDFromHex(id)
	if userId != visitorId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	err = u.UserDBOperator.DeleteUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (u *UserController) GetUser(c *gin.Context) {
	userId, err := utils.GetUSerIdFormToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := u.UserDBOperator.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
