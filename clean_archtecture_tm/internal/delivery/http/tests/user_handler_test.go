package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"task_managemet_api/cmd/task_manager/internal/domain"
	"task_managemet_api/cmd/task_manager/pkg/security"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"

	myHttp "task_managemet_api/cmd/task_manager/internal/delivery/http"
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

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) AddUser(user *domain.User) (*domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserUsecase) CreateAdmin(admin *domain.User) (*domain.User, error) {
	args := m.Called(admin)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserUsecase) DeleteUser(deleterID primitive.ObjectID, tobeDeletedID primitive.ObjectID) error {
	args := m.Called(deleterID, tobeDeletedID)
	return args.Error(0)
}

func (m *MockUserUsecase) UpdateUser(id primitive.ObjectID, user *domain.User) (*domain.User, error) {
	args := m.Called(id, user)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserUsecase) GetUser(id primitive.ObjectID) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserUsecase) LoginUser(email string, password string) (string, string, error) {
	args := m.Called(email, password)
	return args.Get(0).(string), args.Get(1).(string), args.Error(2)
}

func (m *MockUserUsecase) ActivateUser(id primitive.ObjectID) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserUsecase) DeactivateUser(id primitive.ObjectID) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}

type UserHandlerSuit struct {
	suite.Suite
	userUsecase   *MockUserUsecase
	handler       *myHttp.UserHandler
	testingServer *httptest.Server
	// router        *gin.Engine
}

func (suite *UserHandlerSuit) SetupTest() {

	suite.userUsecase = new(MockUserUsecase)
	fromTokenGetter := security.GetTokenData{}

	suite.handler = myHttp.NewUserHandler(suite.userUsecase, fromTokenGetter)

	router := gin.Default()
	r := router.Group("/auth")
	r.POST("/register", suite.handler.AddUser)
	r.POST("/login", suite.handler.LoginUser)
	r.PUT("/user", MockAuthMiddleware(), suite.handler.UpdateUser)
	r.DELETE("/user", MockAuthMiddleware(), suite.handler.DeleteUser)
	r.GET("/user", MockAuthMiddleware(), suite.handler.GetUser)

	admin := router.Group("/admin", MockAuthMiddleware(), MockISAdminMiddleWare())
	admin.PUT("/approve/:id", suite.handler.ApproveUser)
	admin.PUT("/disapprove/:id", suite.handler.DisApproveUser)
	admin.POST("/create-admin", suite.handler.CreateAdmin)

	suite.testingServer = httptest.NewServer(router)

}

func (s *UserHandlerSuit) TearDownSuite() {
	defer s.testingServer.Close()
}

func (s *UserHandlerSuit) TestAddUser_positive() {

	user := &domain.User{
		Email:    "test_user@gmail.com",
		Password: "test_passwordAbcd123@#$",
	}
	s.userUsecase.On("AddUser", mock.Anything).Return(user, nil)

	//marshalling
	requestBody, err := json.Marshal(&user)
	s.NoError(err, "can not marshal struct to json")

	// calling the testing server given the provided request body
	response, err := http.Post(fmt.Sprintf("%s/auth/register", s.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	type responseBody struct {
		Message string      `json:"message"`
		Data    domain.User `json:"data"`
	}

	// Decode the response body
	var resp responseBody
	err = json.NewDecoder(response.Body).Decode(&resp)
	s.NoError(err, "can not unmarshal response body")

	// Run assertions to ensure the correct behavior
	s.Equal(http.StatusOK, response.StatusCode, "expected status code %d got %d", http.StatusOK, response.StatusCode)
	s.Equal("user created successfully", resp.Message, "expected status code %d got %d", http.StatusOK, response.StatusCode)
	s.Equal(user.Email, resp.Data.Email, "expected status code %d got %d", http.StatusOK, response.StatusCode)
	s.userUsecase.AssertExpectations(s.T())

}

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

// func (s *UserHandlerSuit) TestUserUpdate_UseCaseReturnError(t *testing.T) {
// 	user := &domain.User{
// 		Email:    "test_user@gmail.com",
// 		Password: "test_passwordAbcd123@#$",
// 	}
// 	s.userUsecase.On("UpdateUser", mock.Anything, mock.Anything).Return(user, nil)

// 	// Marshalling
// 	requestBody, err := json.Marshal(user)
// 	s.NoError(err, "can not marshal struct to json")

// 	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/auth/user", s.testingServer.URL), bytes.NewBuffer(requestBody))
// 	s.NoError(err, "can not create request")
// 	req.Header.Set("Content-Type", "application/json")

// 	response, err := http.DefaultClient.Do(req)
// 	s.NoError(err, "no error when calling the endpoint")
// 	defer response.Body.Close()

// 	s.Equal(http.StatusOK, response.StatusCode, "expected status code %d got %d", http.StatusOK, response.StatusCode)
// 	s.userUsecase.AssertExpectations(s.T())
// }

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerSuit))
}
