package handler

import (
	"glossika-assignment/internal/middleware"
	"glossika-assignment/internal/model"
	"glossika-assignment/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RecommendationHandler 定義推薦功能的路由處理器
type RecommendationHandler struct {
	service service.IRecommendationService
}

// NewRecommendationHandler 創建新的推薦功能處理器
func NewRecommendationHandler(service service.IRecommendationService) *RecommendationHandler {
	return &RecommendationHandler{
		service: service,
	}
}

// GetRecommendations 獲取推薦內容
func (h *RecommendationHandler) GetRecommendations(c *gin.Context) {
	var req model.RecommendationRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid query parameters: " + err.Error(),
		})
		return
	}

	// 預設值處理
	if req.Limit <= 0 {
		req.Limit = 10 // 預設限制為10條記錄
	}
	if req.Offset < 0 {
		req.Offset = 0 // 預設偏移量為0
	}

	// 獲取推薦項目
	items, err := h.service.GetRecommendations(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	// 返回響應
	c.JSON(http.StatusOK, items)
}

// SetupRoutes 設置推薦相關路由
func (h *RecommendationHandler) SetupRoutes(router *gin.RouterGroup) {
	recRoutes := router.Group("/recommendations")
	recRoutes.Use(middleware.AuthMiddleware()) // 應用身份驗證中間件
	{
		recRoutes.GET("", h.GetRecommendations) // GET 用於透過查詢參數獲取數據
	}
}
