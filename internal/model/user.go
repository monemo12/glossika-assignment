package model

import (
	"time"
)

// User 定義用戶實體
type User struct {
	ID                string    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Email             string    `json:"email"`
	Password          string    `json:"password,omitempty" gorm:"column:password_hash"`
	Name              string    `json:"name"`
	Verified          bool      `json:"verified"`
	VerificationToken string    `json:"verificationToken"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

// RegisterRequest 定義註冊請求
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// RegisterResponse 定義註冊響應
type RegisterResponse struct {
	UserID            string    `json:"userId"`
	Email             string    `json:"email"`
	VerificationToken string    `json:"verificationToken"`
	CreatedAt         time.Time `json:"createdAt"`
}

// LoginRequest 定義登錄請求
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse 定義登錄響應
type LoginResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}

// VerifyEmailRequest 定義驗證郵件請求
type VerifyEmailRequest struct {
	Token string `json:"token"`
}
