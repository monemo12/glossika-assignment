package service

import (
	"context"
	"fmt"
)

// EmailService 定義郵件服務接口
type IEmailService interface {
	SendVerificationEmail(ctx context.Context, email string, token string) error
}

// EmailService 實現郵件服務接口
type EmailService struct {
}

// NewEmailService 創建新的郵件服務
func NewEmailService() *EmailService {
	return &EmailService{}
}

// SendVerificationEmail 發送驗證郵件 (Dummy)
func (s *EmailService) SendVerificationEmail(ctx context.Context, email string, token string) error {
	fmt.Println("Sending verification email to:", email)
	fmt.Println("Token:", token)
	return nil
}
