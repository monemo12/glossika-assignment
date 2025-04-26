package utils

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "Valid password",
			password: "Test123!",
			wantErr:  false,
		},
		{
			name:     "Empty password",
			password: "",
			wantErr:  false, // bcrypt accepts empty strings
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := HashPassword(tt.password)

			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && hash == "" {
				t.Error("HashPassword() returned empty hash for valid input")
			}

			// Verify that the hash works for verification
			if !tt.wantErr {
				err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(tt.password))
				if err != nil {
					t.Errorf("Failed to verify hashed password: %v", err)
				}
			}
		})
	}
}

func TestIsValidPassword(t *testing.T) {
	tests := []struct {
		name           string
		password       string
		wantValid      bool
		wantErrMessage string
	}{
		{
			name:           "Valid password",
			password:       "Test123!",
			wantValid:      true,
			wantErrMessage: "",
		},
		{
			name:           "Too short password",
			password:       "Ab1!",
			wantValid:      false,
			wantErrMessage: "密碼長度必須在 6 到 16 個字元之間",
		},
		{
			name:           "Too long password",
			password:       "Abcdefghijklmnopq1!",
			wantValid:      false,
			wantErrMessage: "密碼長度必須在 6 到 16 個字元之間",
		},
		{
			name:           "No uppercase letter",
			password:       "test123!",
			wantValid:      false,
			wantErrMessage: "密碼必須包含至少一個大寫字母",
		},
		{
			name:           "No lowercase letter",
			password:       "TEST123!",
			wantValid:      false,
			wantErrMessage: "密碼必須包含至少一個小寫字母",
		},
		{
			name:           "No special character",
			password:       "Test1234",
			wantValid:      false,
			wantErrMessage: "密碼必須包含至少一個特殊符號",
		},
		{
			name:           "Has all required characters",
			password:       "Password123!@#",
			wantValid:      true,
			wantErrMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, errMsg := IsValidPassword(tt.password)

			if valid != tt.wantValid {
				t.Errorf("IsValidPassword() valid = %v, wantValid %v", valid, tt.wantValid)
			}

			if errMsg != tt.wantErrMessage {
				t.Errorf("IsValidPassword() errMsg = %v, wantErrMessage %v", errMsg, tt.wantErrMessage)
			}
		})
	}
}
