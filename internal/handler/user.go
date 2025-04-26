package handler

import (
	"glossika-assignment/internal/database"
	"glossika-assignment/internal/model"
	"glossika-assignment/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserHandler 定義用戶功能的路由處理器
type UserHandler struct {
	db database.Database
}

// NewUserHandler 創建新的用戶功能處理器
func NewUserHandler(db database.Database) *UserHandler {
	return &UserHandler{
		db: db,
	}
}

// Register 處理用戶註冊
func (h *UserHandler) Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request data: " + err.Error(),
		})
		return
	}

	// 檢查請求數據
	if req.Email == "" || req.Password == "" || req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Email, password and name are required",
		})
		return
	}

	// 生成用戶 ID (在真實實現中會儲存到數據庫)
	userId := uuid.New().String()
	now := time.Now()

	// 返回響應
	c.JSON(http.StatusCreated, model.RegisterResponse{
		UserID:    userId,
		Email:     req.Email,
		CreatedAt: now,
	})
}

// Login 處理用戶登錄
func (h *UserHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request data: " + err.Error(),
		})
		return
	}

	// 檢查請求數據
	if req.Email == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Email and password are required",
		})
		return
	}

	// 模擬用戶驗證 (在真實實現中會檢查數據庫)
	userId := uuid.New().String()
	now := time.Now()

	// 生成 JWT 令牌，使用配置中的過期時間
	token, expiresAt, err := utils.GenerateToken(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "Failed to generate token: " + err.Error(),
		})
		return
	}

	user := &model.User{
		ID:        userId,
		Email:     req.Email,
		Name:      "Dummy User",
		Verified:  true,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// 返回響應
	c.JSON(http.StatusOK, model.LoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		User:      user,
	})
}

// VerifyEmail 處理用戶驗證
func (h *UserHandler) VerifyEmail(c *gin.Context) {
	var req model.VerifyEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid request data: " + err.Error(),
		})
		return
	}

	// 檢查請求數據
	if req.Token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Token is required",
		})
		return
	}

	// 返回響應
	c.JSON(http.StatusOK, model.VerifyEmailResponse{
		Result: true,
	})
}

// SetupRoutes 設置用戶相關路由
func (h *UserHandler) SetupRoutes(router *gin.RouterGroup) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", h.Register)
		userRoutes.POST("/login", h.Login)
		userRoutes.POST("/verify-email", h.VerifyEmail)
	}
}
