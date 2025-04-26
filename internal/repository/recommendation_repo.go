package repository

import (
	"context"
	"glossika-assignment/internal/model"
	"time"
)

// RecommendationRepository 定義推薦數據訪問接口
type RecommendationRepository interface {
	FetchFromDB(ctx context.Context, userID string, limit, offset int) ([]*model.RecommendationItem, error)
	CacheToRedis(ctx context.Context, userID string, items []*model.RecommendationItem, ttl time.Duration) error
	FetchFromRedis(ctx context.Context, userID string) ([]*model.RecommendationItem, error)
}
