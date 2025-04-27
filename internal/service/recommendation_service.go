package service

import (
	"context"
	"glossika-assignment/internal/model"
	"glossika-assignment/internal/repository"
)

// RecommendationService 定義推薦服務接口
type IRecommendationService interface {
	GetRecommendations(ctx context.Context, req *model.RecommendationRequest) (*model.RecommendationResponse, error)
}

// RecommendationService 實現推薦服務接口
type RecommendationService struct {
	repo repository.IRecommendationRepository
}

// NewRecommendationService 創建新的推薦服務
func NewRecommendationService(repo repository.IRecommendationRepository) *RecommendationService {
	return &RecommendationService{
		repo: repo,
	}
}

// GetRecommendations 獲取推薦項目
func (s *RecommendationService) GetRecommendations(ctx context.Context, req *model.RecommendationRequest) (*model.RecommendationResponse, error) {
	items, err := s.repo.FetchItems(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	return &model.RecommendationResponse{
		Items: items,
	}, nil
}
