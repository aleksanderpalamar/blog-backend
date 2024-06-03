# API Documentation

## Base URL

```bash
http://localhost:8080
```
Endpoints
1. List All Posts
Endpoint: `/posts`
Method: `GET`
Description: Retrieve a list of all blog posts.

Response:

- 200 OK: Returns an array of posts.
```json
{
    "data": [
        {
            "ID": 1,
            "CreatedAt": "2024-05-29T12:34:56Z",
            "UpdatedAt": "2024-05-29T12:34:56Z",
            "DeletedAt": null,
            "title": "First Post",
            "content": ["This is the first post."]
        },
    ]
}
```
3. Create a New Post
Endpoint: `/posts`
Method: `POST`
Description: Create a new blog post.

Request Body:

- `title` (string): The title of the post.
- `content` (array of strings): The content of the post.
**Example Request**
```json
  {
    "title": "New Post",
    "content": ["This is the content of the new post."]
  }
```
**Response**
```json
{
    "data": {
        "ID": 2,
        "CreatedAt": "2024-05-29T13:45:12Z",
        "UpdatedAt": "2024-05-29T13:45:12Z",
        "DeletedAt": null,
        "title": "New Post",
        "content": ["This is the content of the new post."]
    }
}
```
4. Update a Post
Endpoint: `/posts/:id`
Method: `PUT`
Description: Update an existing blog post by its ID.

Path Parameters:

- `id` (integer): The ID of the post to update.
**Request Body:**

- `title` (string): The updated title of the post.
- `content` (array of strings): The updated content of the post.

**Example Request**
```json
{
    "title": "Updated Post",
    "content": ["This is the updated content of the post."]
}
```
Response:

`200 OK`: Returns the updated post.
`400 Bad Request`: Post not found.

```json
{
    "data": {
        "ID": 1,
        "CreatedAt": "2024-05-29T12:34:56Z",
        "UpdatedAt": "2024-05-29T14:56:34Z",
        "DeletedAt": null,
        "title": "Updated Post",
        "content": ["This is the updated content of the post."]
    }
}
```
5. Delete a Post
Endpoint: `/posts/:id`
Method: `DELETE`
Description: Delete a specific blog post by its ID.

Path Parameters:

- `id` (integer): The ID of the post to delete.
Response:

- `200 OK`: Returns a success message.
- `400 Bad Request`: Post not found.
```json
{
  "data": true
}
```
**Error Handling**
All endpoints will return a JSON error response in the following format if an error occurs:
```json
{
    "error": "Error message here"
}
```
## Notes
- Ensure the server is running on `http://localhost:8080`.
- Use appropriate HTTP methods (GET, POST, PUT, DELETE) for each endpoint.
- Make sure to handle all required fields and data types in requests.
This documentation provides a comprehensive overview of the available endpoints and their usage for the blog backend project. If you have any questions or need further assistance, feel free to ask!
