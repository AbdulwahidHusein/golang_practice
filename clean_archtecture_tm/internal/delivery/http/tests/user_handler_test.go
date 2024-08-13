package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"task_managemet_api/cmd/task_manager/internal/domain"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
type UserUseCaseInterface interface {
	AddUser(user *domain.User) (*domain.User, error)
	CreateAdmin(admin *domain.User) (*domain.User, error)
	DeleteUser(deleterID primitive.ObjectID, tobeDeletedID primitive.ObjectID) error
	UpdateUser(id primitive.ObjectID, user *domain.User) (*domain.User, error)
	GetUser(id primitive.ObjectID) (*domain.User, error)
	LoginUser(email string, password string) (string, string, error)
	ActivateUser(id primitive.ObjectID) (*domain.User, error)
	DeactivateUser(id primitive.ObjectID) (*domain.User, error)
}
*/

func (s *UserHandlerSuit) TestAddUser_NoEmail() {
	user := &domain.User{
		Email:    "",
		Password: "test_passwordAbcd123@#$",
	}
	requestBody, err := json.Marshal(&user)
	s.NoError(err, "can not marshal struct to json")

	// calling the testing server given the provided request body
	response, err := http.Post(fmt.Sprintf("%s/auth/register", s.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	s.Equal(http.StatusBadRequest, response.StatusCode, "expected status code %d got %d", http.StatusBadRequest, response.StatusCode)
	s.userUsecase.AssertExpectations(s.T())
}

func (s *UserHandlerSuit) TestAddUser_UseCaseReturnError() {

	user := &domain.User{
		Email:    "test_user@gmail",
		Password: "test_passwordAbcd123@#$",
	}
	s.userUsecase.On("AddUser", user).Return(user, errors.New("some error"))

	//marshalling
	requestBody, err := json.Marshal(user)
	s.NoError(err, "can not marshal struct to json")

	response, err := http.Post(fmt.Sprintf("%s/auth/register", s.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	s.Equal(http.StatusBadRequest, response.StatusCode, "expected status code %d got %d", http.StatusBadRequest, response.StatusCode)
	s.userUsecase.AssertExpectations(s.T())
}

func (s *UserHandlerSuit) TestLoginUser_positive() {
	user := &domain.User{
		Email:    "test_user@gmail.com",
		Password: "test_passwordAbcd123@#$",
	}
	s.userUsecase.On("LoginUser", mock.Anything, mock.Anything).Return("access_token", "refresh_token", nil)

	// Marshalling
	requestBody, err := json.Marshal(user)
	s.NoError(err, "can not marshal struct to json")

	// Calling the testing server given the provided request body
	response, err := http.Post(fmt.Sprintf("%s/auth/login", s.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	var ResponseBody struct {
		Message string `json:"message"`
		Data    struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		} `json:"data"`
	}

	// Decode the response body
	err = json.NewDecoder(response.Body).Decode(&ResponseBody)
	s.NoError(err, "can not unmarshal response body")
	s.Equal(http.StatusOK, response.StatusCode, "expected status code %d got %d", http.StatusOK, response.StatusCode)
	s.Equal("logged in successfully", ResponseBody.Message, "expected message 'logged in successfully', got '%s'", ResponseBody.Message)
	s.Equal("access_token", ResponseBody.Data.AccessToken, "expected access_token 'access_token', got '%s'", ResponseBody.Data.AccessToken)
	s.Equal("refresh_token", ResponseBody.Data.RefreshToken, "expected refresh_token 'refresh_token', got '%s'", ResponseBody.Data.RefreshToken)

	s.userUsecase.AssertExpectations(s.T())
}

func (s *UserHandlerSuit) TestLoginUser_CredentialAsParam() {

	user := &domain.User{
		Email:    "test_user@gmail.com",
		Password: "test_passwordAbcd123@#$",
	}
	// s.userUsecase.On("LoginUser", mock.Anything, mock.Anything).Return("access_token", "refresh_token", nil)

	// // Marshalling

	response, err := http.Post(fmt.Sprintf("%s/auth/login?email=%s&password=%s", s.testingServer.URL, user.Email, user.Password), "application/json", bytes.NewBuffer([]byte{}))
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	s.Equal(http.StatusInternalServerError, response.StatusCode, "expected status code %d got %d", http.StatusBadRequest, response.StatusCode)
	s.userUsecase.AssertExpectations(s.T())
}

func (s *UserHandlerSuit) TestLoginUser_UseCaseReturnError() {

	user := &domain.User{
		Email:    "test_user@gmail.com",
		Password: "test_passwordAbcd123@#$",
	}
	s.userUsecase.On("LoginUser", user.Email, user.Password).Return("", "", errors.New("some error"))

	// Marshalling
	requestBody, err := json.Marshal(user)
	s.NoError(err, "can not marshal struct to json")

	// Calling the testing server given the provided request body
	response, err := http.Post(fmt.Sprintf("%s/auth/login", s.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))

	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()
	s.Equal(http.StatusBadRequest, response.StatusCode, "expected status code %d got %d", http.StatusBadRequest, response.StatusCode)
}

func (s *UserHandlerSuit) TestUpdateUser_Positive() {

	user := &domain.User{
		Email:    "test_user@gmail.com",
		Password: "test_passwordAbcd123@#$",
	}
	s.userUsecase.On("UpdateUser", mock.Anything, mock.Anything).Return(user, nil)

	// Marshalling
	requestBody, err := json.Marshal(user)
	s.NoError(err, "can not marshal struct to json")

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/auth/user", s.testingServer.URL), bytes.NewBuffer(requestBody))
	s.NoError(err, "can not create request")
	req.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	s.Equal(http.StatusOK, response.StatusCode, "expected status code %d got %d", http.StatusOK, response.StatusCode)
	s.userUsecase.AssertExpectations(s.T())
}

func (s *UserHandlerSuit) TestPromoteUser_Positive() {

	user := &domain.User{
		Email:    "test_user@gmail.com",
		Password: "test_passwordAbcd123@#$",
		ID:       primitive.NewObjectID(),
		Role:     "user",
	}
	s.userUsecase.On("PromoteUser", mock.Anything, mock.Anything).Return(user, nil)

	// Marshalling
	requestBody, err := json.Marshal(user)
	s.NoError(err, "can not marshal struct to json")

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/promote/%s", s.testingServer.URL, user.ID.Hex()), bytes.NewBuffer(requestBody))
	s.NoError(err, "can not create request")
	req.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	s.Equal(http.StatusOK, response.StatusCode, "expected status code %d got %d", http.StatusOK, response.StatusCode)
	s.Equal("admin", user.Role, "expected status code %d got %d", http.StatusOK, response.StatusCode)
	s.userUsecase.AssertExpectations(s.T())
}

func (s *UserHandlerSuit) TestDemote_Positive() {
	user := &domain.User{
		Email:    "test_user@gmail.com",
		Password: "test_passwordAbcd123@#$",
		ID:       primitive.NewObjectID(),
		Role:     "user",
	}
	s.userUsecase.On("DemoteUser", mock.Anything, mock.Anything).Return(user, nil)

	// Marshalling
	requestBody, err := json.Marshal(user)
	s.NoError(err, "can not marshal struct to json")

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/demote/%s", s.testingServer.URL, user.ID.Hex()), bytes.NewBuffer(requestBody))
	s.NoError(err, "can not create request")
	req.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	s.Equal(http.StatusOK, response.StatusCode, "expected status code %d got %d", http.StatusOK, response.StatusCode)
	s.Equal("user", user.Role, "expected status code %d got %d", http.StatusOK, response.StatusCode)
	s.userUsecase.AssertExpectations(s.T())
}
func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerSuit))
}
