# API Documentation

## Base URL

`http://<your-server-address>`

## Routes

### Books

#### Add a Book

- **Endpoint**: `/books`
- **Method**: `POST`
- **Description**: Adds a new book to the library.
- **Request Body**: JSON object representing the book (e.g., `{"title": "Book Title", "author": "Author Name", "published_year": 2024}`).
- **Responses**:
  - **201 Created**: Book successfully added.
  - **400 Bad Request**: Invalid request payload.
  - **500 Internal Server Error**: Server error while adding the book.

#### Get All Books

- **Endpoint**: `/books`
- **Method**: `GET`
- **Description**: Retrieves a list of all books in the library.
- **Responses**:
  - **200 OK**: List of books.
  - **500 Internal Server Error**: Error retrieving books.

#### Get a Single Book

- **Endpoint**: `/books/{id}`
- **Method**: `GET`
- **Description**: Retrieves details of a specific book by its ID.
- **Path Parameters**:
  - `id` (string): The unique identifier of the book.
- **Responses**:
  - **200 OK**: Details of the book.
  - **400 Bad Request**: Invalid book ID format.
  - **404 Not Found**: Book not found.
  - **500 Internal Server Error**: Error retrieving the book.

#### Update a Book

- **Endpoint**: `/books/{id}`
- **Method**: `PUT`
- **Description**: Updates details of a specific book by its ID.
- **Path Parameters**:
  - `id` (string): The unique identifier of the book.
- **Request Body**: JSON object with updated book details.
- **Responses**:
  - **200 OK**: Book successfully updated.
  - **400 Bad Request**: Invalid request payload.
  - **404 Not Found**: Book not found.
  - **500 Internal Server Error**: Error updating the book.

#### Delete a Book

- **Endpoint**: `/books/{id}`
- **Method**: `DELETE`
- **Description**: Deletes a specific book by its ID.
- **Path Parameters**:
  - `id` (string): The unique identifier of the book.
- **Responses**:
  - **204 No Content**: Book successfully deleted.
  - **400 Bad Request**: Invalid book ID format.
  - **404 Not Found**: Book not found.
  - **500 Internal Server Error**: Error deleting the book.

#### Borrow a Book

- **Endpoint**: `/borrow`
- **Method**: `POST`
- **Description**: Records a book borrowing transaction.
- **Request Body**: JSON object representing the borrowing details (e.g., `{"book_id": "bookId", "user_id": "userId"}`).
- **Responses**:
  - **200 OK**: Book successfully borrowed.
  - **400 Bad Request**: Invalid request payload or borrowing details.
  - **500 Internal Server Error**: Error processing the borrowing request.

#### Return a Book

- **Endpoint**: `/return`
- **Method**: `POST`
- **Description**: Records a book return transaction.
- **Request Body**: JSON object representing the return details (e.g., `{"book_id": "bookId", "user_id": "userId"}`).
- **Responses**:
  - **200 OK**: Book successfully returned.
  - **400 Bad Request**: Invalid request payload or return details.
  - **500 Internal Server Error**: Error processing the return request.

#### Get Available Books

- **Endpoint**: `/available`
- **Method**: `GET`
- **Description**: Retrieves a list of available books in the library.
- **Responses**:
  - **200 OK**: List of available books.
  - **500 Internal Server Error**: Error retrieving available books.

#### Get Borrowed Books for a User

- **Endpoint**: `/borrowed/{id}`
- **Method**: `GET`
- **Description**: Retrieves a list of borrowed books for a specific user.
- **Path Parameters**:
  - `id` (string): The unique identifier of the user.
- **Responses**:
  - **200 OK**: List of borrowed books.
  - **400 Bad Request**: Invalid user ID format.
  - **404 Not Found**: No borrowed books found for the user.
  - **500 Internal Server Error**: Error retrieving borrowed books.

### Users

#### Add a User

- **Endpoint**: `/users`
- **Method**: `POST`
- **Description**: Adds a new user to the system.
- **Request Body**: JSON object representing the user (e.g., `{"name": "User Name", "email": "user@example.com"}`).
- **Responses**:
  - **201 Created**: User successfully added.
  - **400 Bad Request**: Invalid request payload.
  - **500 Internal Server Error**: Error adding the user.

#### Get All Users

- **Endpoint**: `/users`
- **Method**: `GET`
- **Description**: Retrieves a list of all users.
- **Responses**:
  - **200 OK**: List of users.
  - **500 Internal Server Error**: Error retrieving users.

#### Get a Single User

- **Endpoint**: `/users/{id}`
- **Method**: `GET`
- **Description**: Retrieves details of a specific user by their ID.
- **Path Parameters**:
  - `id` (string): The unique identifier of the user.
- **Responses**:
  - **200 OK**: Details of the user.
  - **400 Bad Request**: Invalid user ID format.
  - **404 Not Found**: User not found.
  - **500 Internal Server Error**: Error retrieving the user.

#### Update a User

- **Endpoint**: `/users/{id}`
- **Method**: `PUT`
- **Description**: Updates details of a specific user by their ID.
- **Path Parameters**:
  - `id` (string): The unique identifier of the user.
- **Request Body**: JSON object with updated user details.
- **Responses**:
  - **200 OK**: User successfully updated.
  - **400 Bad Request**: Invalid request payload.
  - **404 Not Found**: User not found.
  - **500 Internal Server Error**: Error updating the user.

#### Delete a User

- **Endpoint**: `/users/{id}`
- **Method**: `DELETE`
- **Description**: Deletes a specific user by their ID.
- **Path Parameters**:
  - `id` (string): The unique identifier of the user.
- **Responses**:
  - **204 No Content**: User successfully deleted.
  - **400 Bad Request**: Invalid user ID format.
  - **404 Not Found**: User not found.
  - **500 Internal Server Error**: Error deleting the user.

