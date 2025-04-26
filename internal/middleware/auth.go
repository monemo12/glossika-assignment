package middleware

import (
	"glossika-assignment/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 定義身份驗證中間件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 從 Authorization 頭獲取令牌
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "error",
				"error":  "Authorization header is required",
			})
			c.Abort()
			return
		}

		// 檢查 Bearer 格式
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "error",
				"error":  "Authorization header format must be 'Bearer {token}'",
			})
			c.Abort()
			return
		}

		// 解析和驗證令牌
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "error",
				"error":  "Invalid or expired token: " + err.Error(),
			})
			c.Abort()
			return
		}

		// 將用戶 ID 設置到上下文中
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
