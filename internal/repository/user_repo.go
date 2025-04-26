package repository

import (
	"context"
	"glossika-assignment/internal/model"
)

// UserRepository 定義用戶數據訪問接口
type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateUserVerification(ctx context.Context, userID string, verified bool) error
}
