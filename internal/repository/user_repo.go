package repository

import (
	"context"
	"glossika-assignment/internal/database"
	"glossika-assignment/internal/model"

	"gorm.io/gorm"
)

// UserRepository 定義用戶數據訪問接口
type IUserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByVerificationToken(ctx context.Context, token string) (*model.User, error)
	UpdateUserVerification(ctx context.Context, userID string, verified bool) error
	CheckUserExists(ctx context.Context, email string) (bool, error)
}

// UserRepository 實現用戶數據訪問接口
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository 創建新的用戶數據訪問實例
func NewUserRepository(db database.MySQLDatabase) IUserRepository {
	return &UserRepository{
		db: db.GetDB(),
	}
}

// CreateUser 創建用戶
func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// GetUserByEmail 獲取用戶
func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByVerificationToken 獲取用戶
func (r *UserRepository) GetUserByVerificationToken(ctx context.Context, token string) (*model.User, error) {
	var user model.User
	result := r.db.Where("verification_token = ?", token).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUserVerification 更新用戶驗證狀態
func (r *UserRepository) UpdateUserVerification(ctx context.Context, userID string, verified bool) error {
	var user model.User
	result := r.db.Model(&user).Where("id = ?", userID).Update("verified", verified).Update("verification_token", nil)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// CheckUserExists 檢查用戶是否存在
func (r *UserRepository) CheckUserExists(ctx context.Context, email string) (bool, error) {
	var user model.User
	result := r.db.Unscoped().Where("email = ?", email).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		return false, nil
	}

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
