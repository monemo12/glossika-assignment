# User API 文檔

本文檔描述了用戶 API 的使用方法。

## 身份驗證

許多 API 端點需要身份驗證。通過在 HTTP 請求標頭中提供 `Authorization` 標頭完成身份驗證：

```
Authorization: Bearer {your-jwt-token}
```

您可以通過調用登錄 API 獲得 JWT 令牌。

## 端點

所有的 API 端點都以 `/api/v1` 為前綴。

### 用戶註冊

**URL**: `/api/v1/users/register`

**方法**: `POST`

**請求體**:

```json
{
  "email": "user@example.com",
  "password": "yourpassword",
  "name": "User Name"
}
```

**成功響應** (HTTP 201):

```json
{
  "userId": "uuid-string",
  "email": "user@example.com",
  "createdAt": "2023-05-18T10:34:22Z"
}
```

**錯誤響應** (HTTP 400):

```json
{
  "status": "error",
  "error": "錯誤信息"
}
```

### 用戶登錄

**URL**: `/api/v1/users/login`

**方法**: `POST`

**請求體**:

```json
{
  "email": "user@example.com",
  "password": "yourpassword"
}
```

**成功響應** (HTTP 200):

```json
{
  "token": "jwt-token-string",
  "expiresAt": "2023-05-19T10:34:22Z",
  "user": {
    "id": "uuid-string",
    "email": "user@example.com",
    "name": "User Name",
    "verified": true,
    "createdAt": "2023-05-18T10:34:22Z",
    "updatedAt": "2023-05-18T10:34:22Z"
  }
}
```

**錯誤響應** (HTTP 400):

```json
{
  "status": "error",
  "error": "錯誤信息"
}
```

### 驗證郵件

**URL**: `/api/v1/users/verify-email`

**方法**: `POST`

**請求體**:

```json
{
  "token": "verification-token-string"
}
```

**成功響應** (HTTP 200):

```json
{
  "result": true
}
```

**錯誤響應** (HTTP 400):

```json
{
  "status": "error",
  "error": "錯誤信息"
}
```

### 獲取推薦內容

**URL**: `/api/v1/recommendations`

**方法**: `GET`

**需要身份驗證**: 是

**查詢參數**:

- `limit` (可選): 限制返回項目數量，默認為 10
- `offset` (可選): 分頁起始位置，默認為 0

**示例請求**: `/api/v1/recommendations?limit=5&offset=10`

**成功響應** (HTTP 200):

```json
{
  "items": [
    {
      "id": "uuid-string",
      "title": "推薦項目標題",
      "description": "推薦項目的詳細描述"
    },
    // ... 更多項目
  ],
  "total": 15,
  "nextPage": true
}
```

**錯誤響應** (HTTP 400, 401):

```json
{
  "status": "error",
  "error": "錯誤信息"
}
```

## 測試 API

您可以使用 curl 或 Postman 等工具來測試這些 API。

### 註冊示例

```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"yourpassword","name":"User Name"}'
```

### 登錄示例

```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"yourpassword"}'
```

### 獲取推薦示例

```bash
curl -X GET "http://localhost:8080/api/v1/recommendations?limit=5&offset=0" \
  -H "Authorization: Bearer your-jwt-token"
``` 