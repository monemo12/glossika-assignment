package utils

import (
	"fmt"
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword 使用 bcrypt 對密碼進行雜湊
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedBytes), nil
}

// IsValidPassword 檢查密碼是否符合要求：
// 1. 長度在 6 到 16 個字元之間
// 2. 至少包含一個大寫字母
// 3. 至少包含一個小寫字母
// 4. 至少包含一個特殊符號
func IsValidPassword(password string) (bool, string) {
	// 檢查長度
	if len(password) < 6 || len(password) > 16 {
		return false, "密碼長度必須在 6 到 16 個字元之間"
	}

	// 檢查是否包含至少一個大寫字母
	hasUpper := false
	// 檢查是否包含至少一個小寫字母
	hasLower := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsLower(char) {
			hasLower = true
		}
	}

	if !hasUpper {
		return false, "密碼必須包含至少一個大寫字母"
	}

	if !hasLower {
		return false, "密碼必須包含至少一個小寫字母"
	}

	// 檢查是否包含至少一個特殊符號
	specialChars := `\(\)\[\]\{\}<>\+\-\*/\?,\.:;"'_\\\|~\x60!@#\$%\^&=`
	regexPattern := fmt.Sprintf("[%s]", specialChars)
	hasSpecial, err := regexp.MatchString(regexPattern, password)
	if err != nil {
		return false, "密碼驗證錯誤"
	}

	if !hasSpecial {
		return false, "密碼必須包含至少一個特殊符號"
	}

	return true, ""
}
