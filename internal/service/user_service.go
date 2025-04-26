package service

import (
	"context"
	"errors"
	"fmt"
	"glossika-assignment/internal/model"
	"glossika-assignment/internal/repository"
	"glossika-assignment/internal/utils"
	"time"

	"github.com/google/uuid"
)

// UserService 定義用戶服務接口
type IUserService interface {
	Register(ctx context.Context, req *model.RegisterRequest) (*model.RegisterResponse, error)
	Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error)
	VerifyEmail(ctx context.Context, token string) error
}

// UserService 實現用戶服務接口
type UserService struct {
	IUserService
	userRepo     repository.IUserRepository
	emailService IEmailService
}

// NewUserService 創建新的用戶服務
func NewUserService(userRepo repository.IUserRepository, emailService IEmailService) *UserService {
	return &UserService{
		userRepo:     userRepo,
		emailService: emailService,
	}
}

// Register 處理用戶註冊
func (s *UserService) Register(ctx context.Context, req *model.RegisterRequest) (*model.RegisterResponse, error) {
	// 驗證電子郵件
	if valid, message := utils.IsValidEmail(req.Email); !valid {
		return nil, errors.New(message)
	}

	// 檢查密碼是否符合要求
	if valid, message := utils.IsValidPassword(req.Password); !valid {
		return nil, errors.New(message)
	}

	// 檢查是否電子郵件已經存在
	exists, err := s.userRepo.CheckUserExists(ctx, req.Email)
	if err != nil {
		return nil, errors.New("無法檢查電子郵件: " + err.Error())
	}
	if exists {
		return nil, errors.New("該電子郵件已被註冊")
	}

	// 加密密碼
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("密碼加密失敗: " + err.Error())
	}

	// 創建用戶
	now := time.Now()
	verificationToken := uuid.New().String()
	user := &model.User{
		Email:             req.Email,
		Password:          hashedPassword,
		Name:              req.Name,
		VerificationToken: verificationToken,
		CreatedAt:         now,
	}

	// 儲存用戶
	user, err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.New("無法儲存用戶: " + err.Error())
	}

	// 發送驗證郵件
	if err := s.emailService.SendVerificationEmail(ctx, user.Email, verificationToken); err != nil {
		fmt.Println("發送驗證郵件失敗:", err)
	}

	// 返回響應
	return &model.RegisterResponse{
		UserID:            user.ID,
		Email:             user.Email,
		VerificationToken: verificationToken,
		CreatedAt:         user.CreatedAt,
	}, nil
}

// Login 處理用戶登錄
func (s *UserService) Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error) {
	// 驗證電子郵件
	if valid, message := utils.IsValidEmail(req.Email); !valid {
		return nil, errors.New(message)
	}

	// 檢查用戶是否存在
	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		fmt.Println("無法檢查電子郵件: " + err.Error())
		return nil, errors.New("用戶不存在")
	}

	// 檢查用戶是否驗證成功
	if !user.Verified {
		return nil, errors.New("用戶未驗證")
	}

	// 檢查密碼是否正確
	if !utils.CheckPassword(user.Password, req.Password) {
		return nil, errors.New("密碼錯誤")
	}

	// 生成 JWT 令牌
	token, expiresAt, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, errors.New("無法生成 JWT 令牌: " + err.Error())
	}

	return &model.LoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}

// VerifyEmail 處理郵件驗證
func (s *UserService) VerifyEmail(ctx context.Context, token string) error {
	user, err := s.userRepo.GetUserByVerificationToken(ctx, token)
	if err != nil {
		return errors.New("無法驗證郵件: " + err.Error())
	}

	err = s.userRepo.UpdateUserVerification(ctx, user.ID, true)
	if err != nil {
		return errors.New("無法更新用戶驗證狀態: " + err.Error())
	}

	return nil
}
