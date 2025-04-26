package model

import "time"

// User 定義用戶實體
type User struct {
	ID        string
	Email     string
	Password  string
	Name      string
	Verified  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// RegisterRequest 定義註冊請求
type RegisterRequest struct {
	Email    string
	Password string
	Name     string
}

// RegisterResponse 定義註冊響應
type RegisterResponse struct {
	UserID    string
	Email     string
	CreatedAt time.Time
}

// LoginRequest 定義登錄請求
type LoginRequest struct {
	Email    string
	Password string
}

// LoginResponse 定義登錄響應
type LoginResponse struct {
	Token     string
	ExpiresAt time.Time
	User      *User
}
