# Task Management API Documentation

## Overview

This document provides detailed information about the API routes for managing tasks and users in the Task Management API. The API is built using the Gin framework, and all routes that require authentication use JWT tokens. Some routes are restricted to users with admin privileges.

## Task Routes

### **General Task Routes (Requires Authentication)**

These routes are accessible to all authenticated users.

#### **`GET /tasks`**

- **Description**: Retrieves a list of all tasks.
- **Headers**: 
  - `Authorization: Bearer <token>`
- **Sample Request**:
  ```http
  GET /tasks HTTP/1.1
  Host: api.example.com
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
  ```
- **Sample Response**:
  ```json
  [
    {
      "id": "60c72b2f9b7d6c001c8e4ae3",
      "title": "Complete the project",
      "description": "Finalize the remaining tasks",
      "due_date": "2024-08-31T18:30:00Z",
      "status": "pending"
    },
    {
      "id": "60c72b2f9b7d6c001c8e4ae4",
      "title": "Team Meeting",
      "description": "Discuss project updates",
      "due_date": "2024-08-15T09:00:00Z",
      "status": "completed"
    }
  ]
  ```

#### **`GET /tasks/status/:status`**

- **Description**: Retrieves tasks filtered by their status.
- **Headers**: 
  - `Authorization: Bearer <token>`
- **Parameters**: 
  - `:status` - The status of the tasks to filter by (e.g., `pending`, `completed`).
- **Sample Request**:
  ```http
  GET /tasks/status/pending HTTP/1.1
  Host: api.example.com
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
  ```
- **Sample Response**:
  ```json
  [
    {
      "id": "60c72b2f9b7d6c001c8e4ae3",
      "title": "Complete the project",
      "description": "Finalize the remaining tasks",
      "due_date": "2024-08-31T18:30:00Z",
      "status": "pending"
    }
  ]
  ```

#### **`GET /tasks/:id`**

- **Description**: Retrieves the details of a specific task by its ID.
- **Headers**: 
  - `Authorization: Bearer <token>`
- **Parameters**: 
  - `:id` - The ID of the task.
- **Sample Request**:
  ```http
  GET /tasks/60c72b2f9b7d6c001c8e4ae3 HTTP/1.1
  Host: api.example.com
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
  ```
- **Sample Response**:
  ```json
  {
    "id": "60c72b2f9b7d6c001c8e4ae3",
    "title": "Complete the project",
    "description": "Finalize the remaining tasks",
    "due_date": "2024-08-31T18:30:00Z",
    "status": "pending"
  }
  ```

### **Admin Task Routes (Requires Admin Privileges)**

These routes are accessible only to authenticated users with admin privileges. The JWT token provided in the request headers must correspond to a user with an admin role.

#### **`GET /admin/tasks`**

- **Description**: Retrieves a list of all tasks.
- **Headers**: 
  - `Authorization: Bearer <admin-token>`
- **Sample Request**:
  ```http
  GET /admin/tasks HTTP/1.1
  Host: api.example.com
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
  ```
- **Sample Response**:
  ```json
  [
    {
      "id": "60c72b2f9b7d6c001c8e4ae3",
      "title": "Complete the project",
      "description": "Finalize the remaining tasks",
      "due_date": "2024-08-31T18:30:00Z",
      "status": "pending"
    },
    {
      "id": "60c72b2f9b7d6c001c8e4ae4",
      "title": "Team Meeting",
      "description": "Discuss project updates",
      "due_date": "2024-08-15T09:00:00Z",
      "status": "completed"
    }
  ]
  ```

#### **`GET /admin/tasks/:id`**

- **Description**: Retrieves the details of a specific task by its ID.
- **Headers**: 
  - `Authorization: Bearer <admin-token>`
- **Parameters**: 
  - `:id` - The ID of the task.
- **Sample Request**:
  ```http
  GET /admin/tasks/60c72b2f9b7d6c001c8e4ae3 HTTP/1.1
  Host: api.example.com
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
  ```
- **Sample Response**:
  ```json
  {
    "id": "60c72b2f9b7d6c001c8e4ae3",
    "title": "Complete the project",
    "description": "Finalize the remaining tasks",
    "due_date": "2024-08-31T18:30:00Z",
    "status": "pending"
  }
  ```

#### **`PUT /admin/tasks/:id`**

- **Description**: Updates the details of a specific task.
- **Headers**: 
  - `Authorization: Bearer <admin-token>`
- **Parameters**: 
  - `:id` - The ID of the task.
- **Request Body**: JSON containing updated task fields.
- **Sample Request**:
  ```http
  PUT /admin/tasks/60c72b2f9b7d6c001c8e4ae3 HTTP/1.1
  Host: api.example.com
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
  Content-Type: application/json

  {
    "title": "Complete the project quickly",
    "description": "Finalize the remaining tasks with urgency",
    "due_date": "2024-08-25T18:30:00Z",
    "status": "in_progress"
  }
  ```
- **Sample Response**:
  ```json
  {
    "id": "60c72b2f9b7d6c001c8e4ae3",
    "title": "Complete the project quickly",
    "description": "Finalize the remaining tasks with urgency",
    "due_date": "2024-08-25T18:30:00Z",
    "status": "in_progress"
  }
  ```

#### **`POST /admin/tasks`**

- **Description**: Creates a new task.
- **Headers**: 
  - `Authorization: Bearer <admin-token>`
- **Request Body**: JSON containing the new task details.
- **Sample Request**:
  ```http
  POST /admin/tasks HTTP/1.1
  Host: api.example.com
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
  Content-Type: application/json

  {
    "title": "New Task",
    "description": "Create a new task",
    "due_date": "2024-09-01T00:00:00Z",
    "status": "pending"
  }
  ```
- **Sample Response**:
  ```json
  {
    "id": "60c72b2f9b7d6c001c8e4ae5",
    "title": "New Task",
    "description": "Create a new task",
    "due_date": "2024-09-01T00:00:00Z",
    "status": "pending"
  }
  ```

#### **`DELETE /admin/tasks/:id`**

- **Description**: Deletes a specific task by its ID.
- **Headers**: 
  - `Authorization: Bearer <admin-token>`
- **Parameters**: 
  - `:id` - The ID of the task.
- **Sample Request**:
  ```http
  DELETE /admin/tasks/60c72b2f9b7d6c001c8e4ae3 HTTP/1.1
  Host: api.example.com
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
  ```
- **Sample Response**:
  ```json
  {
    "message": "Task deleted successfully"
  }
  ```

### **User Routes (Requires Authentication)**

These routes allow authenticated users to manage their profiles and view task assignments.

#### **`GET /users/`**

- **Description**: Retrieves the authenticated user's profile.
- **Headers**: 
  - `Authorization: Bearer <token>`
- **Sample Request**:
  ```http
  GET /users/me HTTP/1.1
  Host: api.example.com
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
  ```
- **Sample Response**:
  ```json
  {
    "id": "60c72b2f9b7d6c001c8e4ae6",
    "username": "johndoe",
    "email": "johndoe@example.com"
  }
  ```

#### **`PUT /users/`**

- **Description**: Updates the authenticated user's profile.
- **Headers**: 
  - `Authorization: Bearer <token>`
- **Request Body**: JSON containing updated user fields.
- **Sample Request**:
  ```http
  PUT /users/me HTTP/1.1
  Host: api.example.com
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
  Content-Type: application/json

  {
    "email": "newemail@example.com"
  }
  ```
- **Sample Response**:
  ```json
  {
    "id": "60c72b2f9b7d6c001c8e4ae6",
    "username": "johndoe",
    "email": "newemail@example.com"
  }
  ```

### **Admin only Routes (For user management)**

These routes allow admins to control users

#### **`PUT admin/user/approve/:id`**

- **Description**: approves the user changing approved field to True.
- **Headers**: 
  - `Authorization: Bearer <token>`
- **Sample Request**:
  ```http
  PUT /admin/user/approve/60c72b2f9b7d6c001c8e4ae6 HTTP/1.1
  Host: api.example.com
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
  ```
- **Sample Response**:
  ```json
  {
    "id": "60c72b2f9b7d6c001c8e4ae6",
    "username": "johndoe",
    "email": "johndoe@example.com",
    "approved": true
  }
  ```

#### **`PUT admin/user/diapprove/:id`**

- **Description**: dis-approves the user changing approved field to True.
- **Headers**: 
  - `Authorization: Bearer <token>`
- **Sample Request**:
  ```http
  PUT /admin/userr/disapprove/60c72b2f9b7d6c001c8e4ae6 HTTP/1.1
  Host: api.example.com
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
  ```
- **Sample Response**:
  ```json
  {
    "id": "60c72b2f9b7d6c001c8e4ae6",
    "username": "johndoe",
    "email": "johndoe@example.com",
    "approved": false
  }

### **Promote users to admin**

#### **`PUT admin/user/promote/:id`**

- **Description**: promotes a user with the given id to admin.
- **Headers**: 
  - `Authorization: Bearer <token>`
- **Sample Request**:
  ```http
  PUT admin/user/promote/60c72b2f9b7d6c001c8e4ae6 HTTP/1.1
  Host: api.example.com
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
  ```
- **Sample Response**:
  ```json
  {
    "id": "60c72b2f9b7d6c001c8e4ae6",
    "username": "johndoe",
    "email": "johndoe@example.com",
    "approved": true,
    "role" : "admin",
  }
  ```

  ### **demote user from being users to admin**

#### **`PUT admin/user/demote/:id`**

- **Description**: promotes a user with the given id to admin.
- **Headers**: 
  - `Authorization: Bearer <token>`
- **Sample Request**:
  ```http
  PUT /admin/user/demote/60c72b2f9b7d6c001c8e4ae6 HTTP/1.1
  Host: api.example.com
  Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
  ```
- **Sample Response**:
  ```json
  {
    "id": "60c72b2f9b7d6c001c8e4ae6",
    "username": "johndoe",
    "email": "johndoe@example.com",
    "approved": true,
    "role" : "user",
  }

### **Authentication Routes**

These routes are used for user authentication and do not require an existing token.

#### **`POST /auth/login`**

- **Description**: Authenticates a user and returns a JWT token.
- **Request Body**: JSON containing the username and password.
- **Sample Request**:
  ```http
  POST /auth/login HTTP/1.1
  Host: api.example.com
  Content-Type: application/json

  {
    "username": "johndoe",
    "password": "password123"
  }
  ```
- **Sample Response**:
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
  ```

#### **`POST /auth/register`**

- **Description**: Registers a new user and returns a JWT token.
- **Request Body**: JSON containing the new user's details.
- **Sample Request**:
  ```http
  POST /auth/register HTTP/1.1
  Host: api.example.com
  Content-Type: application/json

  {
    "username": "janedoe",
    "email": "janedoe@example.com",
    "password": "password123"
  }
  ```
- **Sample Response**:
  ```json
  {
    "id": "60c72b2f9b7d6c001c8e4ae7",
    "username": "janedoe",
    "email": "janedoe@example.com",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
  ```

## Authentication and Authorization Notes

- **JWT Tokens**: All routes except for the authentication routes (`/auth/login`, `/auth/register`) require a valid JWT token in the `Authorization` header.
- **Admin Routes**: The `/admin` routes require the JWT token to belong to a user with admin privileges. If the token does not correspond to an admin user, the request will be denied with a 403 Forbidden status.
- **User-Specific Routes**: Routes like `/users/me` allow users to manage their own profile. The token must correspond to the user attempting to access or modify the data.

## Error Handling

Common errors include:

- **401 Unauthorized**: When the JWT token is missing or invalid.
- **403 Forbidden**: When attempting to access an admin route without admin privileges.
- **404 Not Found**: When a requested resource (like a task or user) does not exist.

```json
{
  "error": "Unauthorized",
  "message": "Invalid or missing token"
}
```

```json
{
  "error": "Forbidden",
  "message": "You do not have permission to access this resource"
}
```
