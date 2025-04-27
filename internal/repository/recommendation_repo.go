package repository

import (
	"context"
	"errors"
	"time"

	"glossika-assignment/internal/database"
	"glossika-assignment/internal/model"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// RecommendationRepository 定義推薦數據訪問接口
type IRecommendationRepository interface {
	FetchItems(ctx context.Context, limit, offset int) ([]*model.RecommendationItem, error)
	fetchItemsFromDB(ctx context.Context, limit, offset int) ([]*model.RecommendationItem, error)
	fetchItemsFromRedis(ctx context.Context, limit, offset int) ([]*model.RecommendationItem, error)
	cacheItemsToRedis(ctx context.Context, items []*model.RecommendationItem, ttl time.Duration) error
}

type RecommendationRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewRecommendationRepository(db database.MySQLDatabase, redis database.RedisDatabase) *RecommendationRepository {
	return &RecommendationRepository{
		db:  db.GetDB(),
		rdb: redis.GetClient(),
	}
}

func (r *RecommendationRepository) FetchItems(ctx context.Context, limit, offset int) ([]*model.RecommendationItem, error) {
	// 從 Redis 中獲取推薦項目
	items, err := r.fetchItemsFromRedis(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	if len(items) > 0 {
		return items, nil
	}

	// 從 MySQL 中獲取推薦項目
	items, err = r.fetchItemsFromDB(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	// 將推薦項目緩存到 Redis
	err = r.cacheItemsToRedis(ctx, items, 10*time.Second)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *RecommendationRepository) fetchItemsFromDB(ctx context.Context, limit, offset int) ([]*model.RecommendationItem, error) {
	return nil, errors.New("not implemented")
}

func (r *RecommendationRepository) fetchItemsFromRedis(ctx context.Context, limit, offset int) ([]*model.RecommendationItem, error) {
	return nil, errors.New("not implemented")
}

func (r *RecommendationRepository) cacheItemsToRedis(ctx context.Context, items []*model.RecommendationItem, ttl time.Duration) error {
	return errors.New("not implemented")
}
