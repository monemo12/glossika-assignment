package handler

import (
	"glossika-assignment/internal/config"
	"glossika-assignment/internal/database"
	"glossika-assignment/internal/model"
	"glossika-assignment/internal/repository"
	"glossika-assignment/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler 定義用戶功能的路由處理器
type UserHandler struct {
	userService service.IUserService
}

// NewUserHandler 創建新的用戶功能處理器
func NewUserHandler(db database.MySQLDatabase, cfg config.EmailConfig) *UserHandler {

	userRepo := repository.NewUserRepository(db)
	emailService := service.NewEmailService(cfg)
	userService := service.NewUserService(userRepo, emailService)

	return &UserHandler{
		userService: userService,
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

	// 調用服務層執行註冊邏輯
	resp, err := h.userService.Register(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	// 返回響應
	c.JSON(http.StatusCreated, resp)
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

	// 調用服務層執行登錄邏輯
	resp, err := h.userService.Login(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	// 返回響應
	c.JSON(http.StatusOK, resp)
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

	// 調用服務層執行郵件驗證邏輯
	err := h.userService.VerifyEmail(c.Request.Context(), req.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	// 返回響應
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
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
