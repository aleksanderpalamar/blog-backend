# Blog Backend
This project is a backend for a blog application written in GoLang. It provides a complete CRUD (Create, Read, Update, Delete) functionality for managing blog posts. The project uses GORM for ORM and SQLite as the database.

## Project Structure

```bash
/blog
    /main.go
    /routes
        /routes.go
    /models
        /post.go
    /controllers
        /post_controller.go
    /config
        /database.go
        /loadEnvVariables.go
    /middleware
        /corsMiddleware.go
```
## Getting Started
### Prerequisites
- Go installed on your machine.
- SQLite installed (or any other database supported by GORM, but make sure to adjust the connection settings).
- `.env` file with the necessary environment variables (e.g., `PORT`).

**Installing**
Clone the repository:
```bash
git clone https://github.com/aleksanderpalamar/blog-backend.git
cd blog-backend
```
Install dependencies:
```bash
go mod download
```
or
```bash
go mod tidy
```

## Running the application:
```bash
go run main.go
```
The server will start at `http://localhost:8080` (or the port specified in your `.env` file).

### API Endpoints
- GET `/posts`: Retrieve all blog posts.
- GET `/posts/:i`: Retrieve a single blog post by ID.
- POST `/posts`: Create a new blog post.
- PUT `/posts/:id`: Update an existing blog post by ID.
- DELETE `/posts/:id`: Delete a blog post by ID

### API Documentation

[API Documentation](https://github.com/aleksanderpalamar/blog-backend/blob/main/docs/api.md)


### Author: 
[Aleksander Palamar](https://aleksanderpalamar.dev)

### License
This project is licensed under the [MIT License](https://github.com/aleksanderpalamar/blog-backend/blob/main/LICENSE) file for details.