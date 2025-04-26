package utils

import (
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		isValid  bool
		errorMsg string
	}{
		{
			name:     "Valid email",
			email:    "test@example.com",
			isValid:  true,
			errorMsg: "",
		},
		{
			name:     "Valid email with numbers",
			email:    "test123@example.com",
			isValid:  true,
			errorMsg: "",
		},
		{
			name:     "Valid email with special characters",
			email:    "test.user+tag@example-site.co.uk",
			isValid:  true,
			errorMsg: "",
		},
		{
			name:     "Invalid email - no @",
			email:    "testexample.com",
			isValid:  false,
			errorMsg: "無效的電子郵件格式",
		},
		{
			name:     "Invalid email - no domain",
			email:    "test@",
			isValid:  false,
			errorMsg: "無效的電子郵件格式",
		},
		{
			name:     "Invalid email - invalid TLD",
			email:    "test@example.c",
			isValid:  false,
			errorMsg: "無效的電子郵件格式",
		},
		{
			name:     "Invalid email - empty string",
			email:    "",
			isValid:  false,
			errorMsg: "無效的電子郵件格式",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid, errorMsg := IsValidEmail(tt.email)
			if isValid != tt.isValid {
				t.Errorf("IsValidEmail(%q) got valid = %v, want %v", tt.email, isValid, tt.isValid)
			}
			if errorMsg != tt.errorMsg {
				t.Errorf("IsValidEmail(%q) got errorMsg = %q, want %q", tt.email, errorMsg, tt.errorMsg)
			}
		})
	}
}
