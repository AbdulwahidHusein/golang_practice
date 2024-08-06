package http

import (
	"net/http"
	"task_managemet_api/cmd/task_manager/internal/domain"

	"task_managemet_api/cmd/task_manager/pkg/security"
	"task_managemet_api/cmd/task_manager/pkg/validation"

	"task_managemet_api/cmd/task_manager/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUseCase interface {
	AddUser(user *domain.User) error
	DeleteUser(id primitive.ObjectID) error
	UpdateUser(id primitive.ObjectID, user *domain.User) error
	GetUser(id primitive.ObjectID) (*domain.User, error)
	LoginUser(email string) (*domain.User, error)
	CheckUser(email string) (*domain.User, error)
}
type UserHandler struct {
	UserUseCase UserUseCase
}

func NewUserHandler(userUseCase UserUseCase) *UserHandler {
	return &UserHandler{
		UserUseCase: userUseCase,
	}
}

func (u *UserHandler) AddUser(c *gin.Context) {
	var user domain.User
	c.ShouldBindJSON(&user)
	email := user.Email
	password := user.Password

	if !validation.IsValidEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}
	if !validation.IsValidPassword(password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character"})
		return
	}
	if user, _ := u.UserUseCase.CheckUser(email); user != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}
	hashedPassword, _ := security.EncryptPassword(password)
	user.Password = hashedPassword

	err := u.UserUseCase.AddUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (u *UserHandler) LoginUser(c *gin.Context) {
	var guest domain.User
	c.ShouldBindJSON(&guest)
	dbUser, err := u.UserUseCase.LoginUser(guest.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	hashedPassword := dbUser.Password
	if err := security.ComparePassword(hashedPassword, guest.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}
	accessTokenString, refreshTokenString, err := security.CreateToken(dbUser.ID.Hex(), dbUser.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": []map[string]string{{"access": accessTokenString, "refresh": refreshTokenString}}})
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	var user domain.User
	visitorId, err1 := utils.GetUSerIdFormToken(c)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	c.ShouldBindJSON(&user)
	user.ID = visitorId

	err := u.UserUseCase.UpdateUser(user.ID, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
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
	err = u.UserUseCase.DeleteUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (u *UserHandler) GetUser(c *gin.Context) {
	userId, err := utils.GetUSerIdFormToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := u.UserUseCase.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
