
# Library Management System

## Overview

This project is a library management system built with Go. It provides functionalities to manage books and users, including adding, updating, deleting, and retrieving books and user information. The system also supports borrowing and returning books.

## Getting Started

### Prerequisites

- Go 1.18 or higher
- MongoDB for data storage

### Installation

1. **Clone the Repository:**

   ```sh
   git clone https://github.com/AbdulwahidHusein/golang_practice/tree/main/library_management
   cd golang_practice
   cd library_management

Install Dependencies:
go mod tidy

Set Up Environment:
Ensure MongoDB is running and configure the connection settings in your application.

Running the Application
To start the server, use the following command:

go run main.go
The server will start and listen on the specified port (default is 8080).

API Documentation

you can get comprehensive documentation under /doc/APIDoc
Base URL
http://localhost:8080

Routes
Books
Add a Book

Endpoint: /books
Method: POST
Description: Adds a new book to the library.
Request Body: JSON object with book details.
Responses:
201 Created: Successfully added.
400 Bad Request: Invalid request payload.
500 Internal Server Error: Server error.
Get All Books

Endpoint: /books
Method: GET
Description: Retrieves a list of all books.
Responses:
200 OK: List of books.
500 Internal Server Error: Error retrieving books.
Get a Single Book

Endpoint: /books/{id}
Method: GET
Description: Retrieves details of a specific book by its ID.
Path Parameters:
id: The book ID.
Responses:
200 OK: Book details.
400 Bad Request: Invalid book ID.
404 Not Found: Book not found.
500 Internal Server Error: Error retrieving the book.
Update a Book

Endpoint: /books/{id}
Method: PUT
Description: Updates details of a specific book.
Path Parameters:
id: The book ID.
Request Body: JSON object with updated book details.
Responses:
200 OK: Successfully updated.
400 Bad Request: Invalid request payload.
404 Not Found: Book not found.
500 Internal Server Error: Error updating the book.
Delete a Book

Endpoint: /books/{id}
Method: DELETE
Description: Deletes a specific book.
Path Parameters:
id: The book ID.
Responses:
204 No Content: Successfully deleted.
400 Bad Request: Invalid book ID.
404 Not Found: Book not found.
500 Internal Server Error: Error deleting the book.
Borrow a Book

Endpoint: /borrow
Method: POST
Description: Records a book borrowing transaction.
Request Body: JSON object with borrowing details.
Responses:
200 OK: Successfully borrowed.
400 Bad Request: Invalid request payload.
500 Internal Server Error: Error processing the borrowing request.
Return a Book

Endpoint: /return
Method: POST
Description: Records a book return transaction.
Request Body: JSON object with return details.
Responses:
200 OK: Successfully returned.
400 Bad Request: Invalid request payload.
500 Internal Server Error: Error processing the return request.
Get Available Books

Endpoint: /available
Method: GET
Description: Retrieves a list of available books.
Responses:
200 OK: List of available books.
500 Internal Server Error: Error retrieving available books.
Get Borrowed Books for a User

Endpoint: /borrowed/{id}
Method: GET
Description: Retrieves a list of borrowed books for a specific user.
Path Parameters:
id: The user ID.
Responses:
200 OK: List of borrowed books.
400 Bad Request: Invalid user ID.
404 Not Found: No borrowed books found.
500 Internal Server Error: Error retrieving borrowed books.
Users
Add a User

Endpoint: /users
Method: POST
Description: Adds a new user to the system.
Request Body: JSON object with user details.
Responses:
201 Created: Successfully added.
400 Bad Request: Invalid request payload.
500 Internal Server Error: Error adding the user.
Get All Users

Endpoint: /users
Method: GET
Description: Retrieves a list of all users.
Responses:
200 OK: List of users.
500 Internal Server Error: Error retrieving users.
Get a Single User

Endpoint: /users/{id}
Method: GET
Description: Retrieves details of a specific user by their ID.
Path Parameters:
id: The user ID.
Responses:
200 OK: User details.
400 Bad Request: Invalid user ID.
404 Not Found: User not found.
500 Internal Server Error: Error retrieving the user.
Update a User

Endpoint: /users/{id}
Method: PUT
Description: Updates details of a specific user.
Path Parameters:
id: The user ID.
Request Body: JSON object with updated user details.
Responses:
200 OK: Successfully updated.
400 Bad Request: Invalid request payload.
404 Not Found: User not found.
500 Internal Server Error: Error updating the user.
Delete a User

Endpoint: /users/{id}
Method: DELETE
Description: Deletes a specific user.
Path Parameters:
id: The user ID.
Responses:
204 No Content: Successfully deleted.
400 Bad Request: Invalid user ID.
404 Not Found: User not found.
500 Internal Server Error: Error deleting the user.
Error Handling
The API returns standard HTTP status codes to indicate success or failure of requests. The response body usually contains an error message describing the issue.

Contributing
Contributions are welcome! Please fork the repository and submit a pull request with your changes.