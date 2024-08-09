package http

import (
	"net/http"
	"task_managemet_api/cmd/task_manager/internal/domain"
	"task_managemet_api/cmd/task_manager/pkg/security"

	"task_managemet_api/cmd/task_manager/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	UserUseCase usecase.UserUseCaseInterface
}

func NewUserHandler(userUseCase usecase.UserUseCaseInterface) *UserHandler {
	return &UserHandler{
		UserUseCase: userUseCase,
	}
}

func (u *UserHandler) AddUser(c *gin.Context) {
	var user domain.User
	c.ShouldBindJSON(&user)

	usr, err := u.UserUseCase.AddUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occured", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user created successfully", "data": usr})
}

func (u *UserHandler) LoginUser(c *gin.Context) {
	var guest domain.User
	c.ShouldBindJSON(&guest)
	accessTokenString, refreshTokenString, err := u.UserUseCase.LoginUser(guest.Email, guest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occured", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "logged in successfully", "data": []map[string]string{{"access": accessTokenString, "refresh": refreshTokenString}}})
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	var user domain.User
	userId, err1 := security.GetUSerIdFormToken(c)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "an error occured", "error": err1.Error()})
		return
	}
	c.ShouldBindJSON(&user)

	updated := u.UserUseCase.UpdateUser(userId, &user)
	if updated == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occured", "error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	deleterID, err := security.GetUSerIdFormToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "an error occured", "error": err.Error()})
		return
	}
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "an error occured", "error": err})
		return
	}
	err = u.UserUseCase.DeleteUser(deleterID, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occured", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully", "data": userId})
}

func (u *UserHandler) GetUser(c *gin.Context) {
	userId, err := security.GetUSerIdFormToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "an invalid id", "error": "Invalid user ID"})
		return
	}
	user, err := u.UserUseCase.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occured", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) ApproveUser(c *gin.Context) {
	userId := c.Param("id")
	userIdObj, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id", "error": "Invalid user ID"})
		return
	}
	user, err1 := u.UserUseCase.ActivateUser(userIdObj)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occured", "error": err1.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User approved successfully", "user": user})
}

func (u *UserHandler) DisApproveUser(c *gin.Context) {
	userId := c.Param("id")
	userIdObj, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id", "error": "Invalid user ID"})
		return
	}
	user, err1 := u.UserUseCase.DeactivateUser(userIdObj)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occured", "error": err1.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User approved successfully", "data": user})
}

func (u *UserHandler) CreateAdmin(c *gin.Context) {
	var user domain.User
	c.ShouldBindJSON(&user)
	user.Role = "admin"
	admin, err := u.UserUseCase.CreateAdmin(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occured", "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": admin})
}
