package service

import (
	"context"
	"glossika-assignment/internal/model"
)

// RecommendationService 定義推薦服務接口
type IRecommendationService interface {
	GetRecommendations(ctx context.Context, req *model.RecommendationRequest) (*model.RecommendationResponse, error)
}
