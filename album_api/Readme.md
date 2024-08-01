
# Web Service - Gin

This is a simple web service built with the Gin framework in Go. It provides a basic API for managing a collection of music albums.

## Project Structure

- `main.go`: The entry point of the application that sets up the routes and starts the server.
- `store/album_store.go`: Contains a mock store for storing album data.
- `models/album.go`: Defines the data structure for an album.
- `controllers/album_controller.go`: Contains the API handlers for managing albums.

## Getting Started

### Prerequisites

- Go installed on your machine (version 1.16 or later).

### Installation

1. **Clone the repository**:
 ```sh
   git clone https://github.com/AbdulwahidHusein/golang_practice/tree/main/web-service-gin
   cd golang_practice
   cd web-services-gin
```

### Install the required Go modules:

``` sh
go mod tidy
```

```sh
go run main.go
```

The server will start and listen on localhost:8080.

API Endpoints
Get All Albums
Endpoint: GET /albums
Description: Retrieves a list of all albums.
Response:
Status Code: 200 OK
Body: Array of album objects.
Create a New Album
Endpoint: POST /albums
Description: Adds a new album to the collection.
Request Body: JSON object representing the album. Example:

```sh
{
  "id": "4",
  "title": "Kind of Blue",
  "artist": "Miles Davis",
  "price": 59.99
}
Response:
Status Code: 201 Created
Body: JSON object representing the created album.
Get an Album by ID
Endpoint: GET /albums/:id
Description: Retrieves a specific album by its ID.
Parameters:
id: The ID of the album to retrieve.
Response:
Status Code: 200 OK if the album is found.
Status Code: 404 Not Found if the album is not found.
Body: JSON object representing the album or an error message.
Example Requests


### Get All Albums


curl -X GET http://localhost:8080/albums
Create a New Album
sh
Copy code
curl -X POST http://localhost:8080/albums -H "Content-Type: application/json" -d '{"id": "4", "title": "Kind of Blue", "artist": "Miles Davis", "price": 59.99}'



### Get an Album by ID

curl -X GET http://localhost:8080/albums/1
Directory Structure
main.go: Entry point for the application.
store/album_store.go: Mock data store for albums.
models/album.go: Data model for albums.
controllers/album_controller.go: API handlers for album-related operations.
Contributing
Feel free to open issues or submit pull requests if you have suggestions or improvements.

```







