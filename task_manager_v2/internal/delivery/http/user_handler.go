package http

import (
	"net/http"
	"task_managemet_api/cmd/task_manager/internal/domain"
	"task_managemet_api/cmd/task_manager/pkg/security"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUseCase interface {
	AddUser(user *domain.User) error
	CreateAdmin(admin *domain.User) error
	DeleteUser(deleterID primitive.ObjectID, tobeDeletedID primitive.ObjectID) error
	UpdateUser(id primitive.ObjectID, user *domain.User) *domain.User
	GetUser(id primitive.ObjectID) (*domain.User, error)
	LoginUser(email string, password string) (string, string, error)
	ApproveUser(id primitive.ObjectID) (*domain.User, error)
	DisApproveUser(id primitive.ObjectID) (*domain.User, error)
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

	updated := u.UserUseCase.AddUser(&user)
	if updated == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (u *UserHandler) LoginUser(c *gin.Context) {
	var guest domain.User
	c.ShouldBindJSON(&guest)
	accessTokenString, refreshTokenString, err := u.UserUseCase.LoginUser(guest.Email, guest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": []map[string]string{{"access": accessTokenString, "refresh": refreshTokenString}}})
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	var user domain.User
	userId, err1 := security.GetUSerIdFormToken(c)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	c.ShouldBindJSON(&user)

	updated := u.UserUseCase.UpdateUser(userId, &user)
	if updated == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	deleterID, err := security.GetUSerIdFormToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	err = u.UserUseCase.DeleteUser(deleterID, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (u *UserHandler) GetUser(c *gin.Context) {
	userId, err := security.GetUSerIdFormToken(c)
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

func (u *UserHandler) ApproveUser(c *gin.Context) {
	userId := c.Param("id")
	userIdObj, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err1 := u.UserUseCase.ApproveUser(userIdObj)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err1.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User approved successfully", "user": user})
}

func (u *UserHandler) DisApproveUser(c *gin.Context) {
	userId := c.Param("id")
	userIdObj, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err1 := u.UserUseCase.ApproveUser(userIdObj)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err1.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User approved successfully", "user": user})
}

func (u *UserHandler) CreateAdmin(c *gin.Context) {
	var user domain.User
	c.ShouldBindJSON(&user)
	user.Role = "admin"
	updated := u.UserUseCase.AddUser(&user)
	if updated == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
