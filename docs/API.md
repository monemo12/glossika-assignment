# API Documentation

All endpoints are prefixed with `/api/v1`.

## Authentication

Some endpoints require authentication via JWT. Provide the token in the HTTP header:

```
Authorization: Bearer {your-jwt-token}
```

---

## User APIs

### Register a New User

- **URL:** `/api/v1/users/register`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "email": "user@example.com",
    "password": "yourpassword",
    "name": "User Name"
  }
  ```
- **Success Response** (`201 Created`):
  ```json
  {
    "userId": "uuid-string",
    "email": "user@example.com",
    "verificationToken": "verification-token-string",
    "createdAt": "2023-05-18T10:34:22Z"
  }
  ```
- **Error Response** (`400 Bad Request`):
  ```json
  {
    "status": "error",
    "error": "Error message"
  }
  ```

---

### User Login

- **URL:** `/api/v1/users/login`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "email": "user@example.com",
    "password": "yourpassword"
  }
  ```
- **Success Response** (`200 OK`):
  ```json
  {
    "token": "jwt-token-string",
    "expiresAt": "2023-05-19T10:34:22Z"
  }
  ```
- **Error Response** (`400 Bad Request`):
  ```json
  {
    "status": "error",
    "error": "Error message"
  }
  ```

---

### Verify Email

- **URL:** `/api/v1/users/verify-email`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "token": "verification-token-string"
  }
  ```
- **Success Response** (`200 OK`):
  ```json
  {
    "status": "success"
  }
  ```
- **Error Response** (`400 Bad Request`):
  ```json
  {
    "status": "error",
    "error": "Error message"
  }
  ```

---

## Recommendation APIs

### Get Recommendations

- **URL:** `/api/v1/recommendations`
- **Method:** `GET`
- **Authentication Required:** Yes
- **Query Parameters:**
  - `limit` (optional, integer): Number of items to return (default: 10)
  - `offset` (optional, integer): Pagination offset (default: 0)
- **Example Request:**
  `/api/v1/recommendations?limit=5&offset=10`
- **Success Response** (`200 OK`):
  ```json
  {
    "items": [
      {
        "id": 1,
        "title": "Recommendation Title",
        "description": "Detailed description of the recommendation.",
        "score": 4.5,
        "created_at": "2023-05-18T10:34:22Z",
        "updated_at": "2023-05-18T10:34:22Z"
      }
      // ... more items
    ],
    "total": 15,
    "nextPage": true
  }
  ```
- **Error Response** (`400 Bad Request`, `401 Unauthorized`):
  ```json
  {
    "status": "error",
    "error": "Error message"
  }
  ```
