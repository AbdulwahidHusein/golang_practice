
### Task Manager API

This repository contains the source code for a Task Manager API built using the Go programming language and the Gin web framework. The API provides endpoints for managing tasks, including creating, retrieving, updating, and deleting tasks, with the data stored in a MongoDB database.

The API also includes endpoints for user authentication and authorization using JWT, ensuring that users can securely perform CRUD operations on their own tasks.

This repository is built in adherence to clean architecture principles. By following these principles, the codebase is structured to promote separation of concerns, making it more manageable and easily scalable. The clean architecture ensures that the business logic is independent of the framework, UI, or external agencies, allowing for flexibility in adapting to changing requirements and facilitating easier testing and maintenance.
folder structure

task_manager_v2/<br/>
├── cmd//<br/>
│   └── taskmanager//<br/>
│       └── main.go/<br/>
├── configs//<br/>
│   └── config.yaml/<br/>
|   └── get_env.go/<br/>
├── internal//<br/>
│   ├── domain//<br/>
│   │   ├── task.go/<br/>
│   │   └── user.go/<br/>
│   ├── repository//<br/>
│   │   ├── task_repository.go/<br/>
│   │   └── user_repository.go/<br/>
│   ├── usecase//<br/>
│   │   ├── task_usecase.go/<br/>
│   │   └── user_usecase.go/<br/>
│   ├── delivery//<br/>
│   │   ├── http//<br/>
│   │   │   ├── task_handler.go/<br/>
│   │   │   └── user_handler.go/<br/>
│   │   └── middleware//<br/>
│   │       └── auth_middleware.go/<br/>
├── pkg//<br/>
│   ├── db//<br/>
│   │   └── mongo.go/<br/>
│   └── security//<br/>
│       └── tokens.go/<br/>
|       └── password_encryption.go/<br/>
|       └── extract.go/<br/>
|   └── validation//<br/>
|       └── user_validation.go/<br/>
├── go.mod/<br/>
|
├── go.sum/<br/>
|
├── Readme.md/<br/>



### Installation
## Clone the repository:

```
git clone https://github.com/abdulwahidHussein/golang_practice.git
```
# Change to the project directory:

```
cd task_manager_api
cd task_manager_api_v2/cmd/task_manager
```

## Install dependencies:

```
go mod tidy
```


### Run the application:

```
go run main.go
```




### API DOCUMENTATION: Visit 
<a href="https://documenter.getpostman.com/view/28093624/2sA3rwMZaX"> Documentation Link</a>