package service

import (
	"context"
	"glossika-assignment/internal/model"
)

// UserService 定義用戶服務接口
type UserService interface {
	Register(ctx context.Context, req *model.RegisterRequest) (*model.RegisterResponse, error)
	Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error)
	VerifyEmail(ctx context.Context, token string) error
}
