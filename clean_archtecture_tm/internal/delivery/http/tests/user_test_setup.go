package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"task_managemet_api/cmd/task_manager/internal/domain"
	"task_managemet_api/cmd/task_manager/pkg/security"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"

	myHttp "task_managemet_api/cmd/task_manager/internal/delivery/http"
)

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
func (m *MockUserUsecase) DemoteUser(id primitive.ObjectID) (*domain.User, error) {
	args := m.Called(id)
	args[0].(*domain.User).Role = "user"
	return args.Get(0).(*domain.User), args.Error(1)
}
func (m *MockUserUsecase) PromoteUser(id primitive.ObjectID) (*domain.User, error) {
	args := m.Called(id)
	args[0].(*domain.User).Role = "admin"
	return args[0].(*domain.User), args.Error(1)
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

	admin.PUT("/promote/:id", suite.handler.PromoteUser)
	admin.PUT("/demote/:id", suite.handler.DemoteUser)

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
