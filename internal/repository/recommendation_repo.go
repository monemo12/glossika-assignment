package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"glossika-assignment/internal/database"
	"glossika-assignment/internal/model"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// RecommendationRepository 定義推薦數據訪問接口
type IRecommendationRepository interface {
	FetchItemsByPagination(ctx context.Context, limit, offset int) ([]*model.Recommendation, error)
	FetchItemsCount(ctx context.Context) (int, error)
	fetchItemsFromDB(ctx context.Context, limit, offset int) ([]*model.Recommendation, error)
	fetchItemsFromRedis(ctx context.Context, limit, offset int) ([]*model.Recommendation, error)
	cacheItemsToRedis(ctx context.Context, items []*model.Recommendation, ttl time.Duration) error
	fetchItemsCountFromDB(ctx context.Context) (int, error)
	fetchItemsCountFromRedis(ctx context.Context) (int, error)
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

const (
	recommendationKey           = "recommendations"
	recommendationPageKeyFormat = "recommendations:page:%d:%d" // format: recommendations:page:limit:offset
	recommendationAllKey        = "recommendations:all"
	cacheBatchSize              = 500 // 緩存批次大小，每次從數據庫多取數據
	defaultCacheTTL             = 10 * time.Minute
)

func (r *RecommendationRepository) FetchItemsByPagination(ctx context.Context, limit, offset int) ([]*model.Recommendation, error) {
	// 從 Redis 中獲取推薦項目
	items, err := r.fetchItemsFromRedis(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	// 如果從 Redis 中獲取到足夠的數據，直接返回
	if len(items) == limit {
		return items, nil
	}

	// 從數據庫獲取當前請求所需的數據
	items, err = r.fetchItemsFromDB(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	// 異步預加載更多數據到緩存
	go func() {
		bgCtx := context.Background()
		// 從數據庫獲取更多數據用於緩存
		cacheItems, err := r.fetchItemsFromDB(bgCtx, cacheBatchSize, 0)
		if err != nil {
			fmt.Printf("Error fetching items for cache: %v\n", err)
			return
		}

		// 緩存數據到Redis
		err = r.cacheItemsToRedis(bgCtx, cacheItems, defaultCacheTTL)
		if err != nil {
			fmt.Printf("Error caching items to Redis: %v\n", err)
		}
	}()

	return items, nil
}

func (r *RecommendationRepository) FetchItemsCount(ctx context.Context) (int, error) {
	count, err := r.fetchItemsCountFromRedis(ctx)
	if err != nil {
		fmt.Println("fetchItemsCountFromRedis error", err)
	}

	if count > 0 {
		return count, nil
	}

	count, err = r.fetchItemsCountFromDB(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *RecommendationRepository) fetchItemsFromDB(ctx context.Context, limit, offset int) ([]*model.Recommendation, error) {
	items := make([]*model.Recommendation, 0, limit)
	err := r.db.WithContext(ctx).Order("score DESC").Limit(limit).Offset(offset).Find(&items).Error
	// 模擬慢速查詢
	time.Sleep(3 * time.Second)
	return items, err
}

func (r *RecommendationRepository) fetchItemsFromRedis(ctx context.Context, limit, offset int) ([]*model.Recommendation, error) {
	// 使用 ZRANGE 從有序集合中獲取指定範圍的項目
	results, err := r.rdb.ZRevRange(ctx, recommendationKey, int64(offset), int64(offset+limit-1)).Result()

	// 如果沒有找到任何結果，返回空切片
	if err != nil {
		if err == redis.Nil {
			return make([]*model.Recommendation, 0), nil
		}
		return nil, err
	}

	// 將結果反序列化為 Recommendation 切片
	items := make([]*model.Recommendation, 0, len(results))
	for _, result := range results {
		var item model.Recommendation
		if err := json.Unmarshal([]byte(result), &item); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}

func (r *RecommendationRepository) cacheItemsToRedis(ctx context.Context, items []*model.Recommendation, ttl time.Duration) error {
	// 在更新前先清除舊的數據
	r.rdb.Del(ctx, recommendationKey)

	// 使用 pipeline 批量添加數據
	pipe := r.rdb.Pipeline()

	for _, item := range items {
		// 將項目序列化為 JSON
		data, err := json.Marshal(item)
		if err != nil {
			return err
		}

		// 添加到有序集合
		pipe.ZAdd(ctx, recommendationKey, redis.Z{
			Score:  item.Score,
			Member: string(data),
		})
	}

	// 設置過期時間
	pipe.Expire(ctx, recommendationKey, ttl)

	// 執行 pipeline
	_, err := pipe.Exec(ctx)
	return err
}

func (r *RecommendationRepository) fetchItemsCountFromDB(ctx context.Context) (int, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Recommendation{}).Count(&count).Error
	return int(count), err
}

func (r *RecommendationRepository) fetchItemsCountFromRedis(ctx context.Context) (int, error) {
	count, err := r.rdb.ZCard(ctx, recommendationKey).Result()
	return int(count), err
}
