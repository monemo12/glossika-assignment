package service

import "context"

// EmailService 定義郵件服務接口
type EmailService interface {
	SendVerificationEmail(ctx context.Context, email string, token string) error
}
