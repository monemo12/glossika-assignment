package service

import (
	"context"
	"testing"
)

func TestSendVerificationEmail(t *testing.T) {
	emailService := NewEmailService()
	ctx := context.Background()
	email := "test@example.com"
	token := "dummy-token"

	err := emailService.SendVerificationEmail(ctx, email, token)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
