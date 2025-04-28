package service

import (
	"context"
	"errors"
	"testing"

	"glossika-assignment/internal/model"
	"glossika-assignment/internal/repository"
	"glossika-assignment/internal/utils"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// MockEmailService creates a mock email service for testing
type MockEmailService struct {
	sendVerificationEmailFunc func(ctx context.Context, email string, token string) error
}

func (m *MockEmailService) SendVerificationEmail(ctx context.Context, email string, token string) error {
	if m.sendVerificationEmailFunc != nil {
		return m.sendVerificationEmailFunc(ctx, email, token)
	}
	return nil
}

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := repository.NewMockIUserRepository(ctrl)
	mockEmailService := &MockEmailService{}

	userService := NewUserService(mockUserRepo, mockEmailService)

	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		// Arrange
		req := &model.RegisterRequest{
			Email:    "test@example.com",
			Password: "Password123!",
			Name:     "Test User",
		}

		mockUserRepo.EXPECT().
			CheckUserExists(gomock.Any(), req.Email).
			Return(false, nil)

		mockUserRepo.EXPECT().
			CreateUser(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ context.Context, user *model.User) (*model.User, error) {
				user.ID = "user-123"
				return user, nil
			})

		mockEmailService.sendVerificationEmailFunc = func(ctx context.Context, email string, token string) error {
			return nil
		}

		// Act
		resp, err := userService.Register(ctx, req)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "user-123", resp.UserID)
		assert.Equal(t, req.Email, resp.Email)
		assert.NotEmpty(t, resp.VerificationToken)
		assert.False(t, resp.CreatedAt.IsZero())
	})

	t.Run("Invalid Email", func(t *testing.T) {
		// Arrange
		req := &model.RegisterRequest{
			Email:    "invalid-email",
			Password: "Password123!",
			Name:     "Test User",
		}

		// Act
		resp, err := userService.Register(ctx, req)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "無效的電子郵件格式")
	})

	t.Run("Invalid Password", func(t *testing.T) {
		// Arrange
		req := &model.RegisterRequest{
			Email:    "test@example.com",
			Password: "weak",
			Name:     "Test User",
		}

		// Act
		resp, err := userService.Register(ctx, req)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "密碼")
	})

	t.Run("Email Already Exists", func(t *testing.T) {
		// Arrange
		req := &model.RegisterRequest{
			Email:    "existing@example.com",
			Password: "Password123!",
			Name:     "Test User",
		}

		mockUserRepo.EXPECT().
			CheckUserExists(gomock.Any(), req.Email).
			Return(true, nil)

		// Act
		resp, err := userService.Register(ctx, req)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "已被註冊")
	})

	t.Run("Repository Error - Check User", func(t *testing.T) {
		// Arrange
		req := &model.RegisterRequest{
			Email:    "test@example.com",
			Password: "Password123!",
			Name:     "Test User",
		}

		mockUserRepo.EXPECT().
			CheckUserExists(gomock.Any(), req.Email).
			Return(false, errors.New("database error"))

		// Act
		resp, err := userService.Register(ctx, req)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "無法檢查電子郵件")
	})

	t.Run("Repository Error - Create User", func(t *testing.T) {
		// Arrange
		req := &model.RegisterRequest{
			Email:    "test@example.com",
			Password: "Password123!",
			Name:     "Test User",
		}

		mockUserRepo.EXPECT().
			CheckUserExists(gomock.Any(), req.Email).
			Return(false, nil)

		mockUserRepo.EXPECT().
			CreateUser(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("database error"))

		// Act
		resp, err := userService.Register(ctx, req)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "無法儲存用戶")
	})
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := repository.NewMockIUserRepository(ctrl)
	mockEmailService := &MockEmailService{}

	userService := NewUserService(mockUserRepo, mockEmailService)

	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		// Arrange
		req := &model.LoginRequest{
			Email:    "test@example.com",
			Password: "Password123!",
		}

		hashedPassword, _ := utils.HashPassword(req.Password)
		user := &model.User{
			ID:       "user-123",
			Email:    req.Email,
			Password: hashedPassword,
			Verified: true,
		}

		mockUserRepo.EXPECT().
			GetUserByEmail(gomock.Any(), req.Email).
			Return(user, nil)

		// Act
		resp, err := userService.Login(ctx, req)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Token)
		assert.False(t, resp.ExpiresAt.IsZero())
	})

	t.Run("Invalid Email", func(t *testing.T) {
		// Arrange
		req := &model.LoginRequest{
			Email:    "invalid-email",
			Password: "Password123!",
		}

		// Act
		resp, err := userService.Login(ctx, req)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "無效的電子郵件格式")
	})

	t.Run("User Not Found", func(t *testing.T) {
		// Arrange
		req := &model.LoginRequest{
			Email:    "test@example.com",
			Password: "Password123!",
		}

		mockUserRepo.EXPECT().
			GetUserByEmail(gomock.Any(), req.Email).
			Return(nil, errors.New("user not found"))

		// Act
		resp, err := userService.Login(ctx, req)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "用戶不存在")
	})

	t.Run("User Not Verified", func(t *testing.T) {
		// Arrange
		req := &model.LoginRequest{
			Email:    "test@example.com",
			Password: "Password123!",
		}

		hashedPassword, _ := utils.HashPassword(req.Password)
		user := &model.User{
			ID:       "user-123",
			Email:    req.Email,
			Password: hashedPassword,
			Verified: false,
		}

		mockUserRepo.EXPECT().
			GetUserByEmail(gomock.Any(), req.Email).
			Return(user, nil)

		// Act
		resp, err := userService.Login(ctx, req)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "用戶未驗證")
	})

	t.Run("Wrong Password", func(t *testing.T) {
		// Arrange
		req := &model.LoginRequest{
			Email:    "test@example.com",
			Password: "WrongPassword123!",
		}

		hashedPassword, _ := utils.HashPassword("Password123!")
		user := &model.User{
			ID:       "user-123",
			Email:    req.Email,
			Password: hashedPassword,
			Verified: true,
		}

		mockUserRepo.EXPECT().
			GetUserByEmail(gomock.Any(), req.Email).
			Return(user, nil)

		// Act
		resp, err := userService.Login(ctx, req)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "密碼錯誤")
	})
}

func TestVerifyEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := repository.NewMockIUserRepository(ctrl)
	mockEmailService := &MockEmailService{}

	userService := NewUserService(mockUserRepo, mockEmailService)

	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		// Arrange
		token := "test-verification-token"
		user := &model.User{
			ID:                "user-123",
			Email:             "test@example.com",
			VerificationToken: token,
			Verified:          false,
		}

		mockUserRepo.EXPECT().
			GetUserByVerificationToken(gomock.Any(), token).
			Return(user, nil)

		mockUserRepo.EXPECT().
			UpdateUserVerification(gomock.Any(), user.ID, true).
			Return(nil)

		// Act
		err := userService.VerifyEmail(ctx, token)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Invalid Token", func(t *testing.T) {
		// Arrange
		token := "invalid-token"

		mockUserRepo.EXPECT().
			GetUserByVerificationToken(gomock.Any(), token).
			Return(nil, errors.New("token not found"))

		// Act
		err := userService.VerifyEmail(ctx, token)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "無法驗證郵件")
	})

	t.Run("Update Error", func(t *testing.T) {
		// Arrange
		token := "test-verification-token"
		user := &model.User{
			ID:                "user-123",
			Email:             "test@example.com",
			VerificationToken: token,
			Verified:          false,
		}

		mockUserRepo.EXPECT().
			GetUserByVerificationToken(gomock.Any(), token).
			Return(user, nil)

		mockUserRepo.EXPECT().
			UpdateUserVerification(gomock.Any(), user.ID, true).
			Return(errors.New("database error"))

		// Act
		err := userService.VerifyEmail(ctx, token)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "無法更新用戶驗證狀態")
	})
}
