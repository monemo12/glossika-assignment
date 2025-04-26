package handler

import (
	"glossika-assignment/internal/database"
	"glossika-assignment/internal/middleware"
	"glossika-assignment/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RecommendationHandler 定義推薦功能的路由處理器
type RecommendationHandler struct {
	db database.IDatabase
}

// NewRecommendationHandler 創建新的推薦功能處理器
func NewRecommendationHandler(db database.IDatabase) *RecommendationHandler {
	return &RecommendationHandler{
		db: db,
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

	// 生成 dummy 數據
	items := make([]*model.RecommendationItem, 0, req.Limit)
	for i := 0; i < req.Limit; i++ {
		items = append(items, &model.RecommendationItem{
			ID:          uuid.New().String(),
			Title:       "推薦項目 " + uuid.New().String()[0:8],
			Description: "這是一個推薦項目的詳細描述",
		})
	}

	// 返回響應
	c.JSON(http.StatusOK, model.RecommendationResponse{
		Items:    items,
		Total:    req.Limit + 5, // 假設還有更多項目
		NextPage: true,
	})
}

// SetupRoutes 設置推薦相關路由
func (h *RecommendationHandler) SetupRoutes(router *gin.RouterGroup) {
	recRoutes := router.Group("/recommendations")
	recRoutes.Use(middleware.AuthMiddleware()) // 應用身份驗證中間件
	{
		recRoutes.GET("", h.GetRecommendations) // GET 用於透過查詢參數獲取數據
	}
}
