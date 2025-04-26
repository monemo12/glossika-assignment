package utils

import "regexp"

// IsValidEmail 檢查電子郵件是否有效
func IsValidEmail(email string) (bool, string) {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return false, "無效的電子郵件格式"
	}
	return true, ""
}
