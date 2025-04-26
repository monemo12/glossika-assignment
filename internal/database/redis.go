package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"glossika-assignment/internal/config"

	"github.com/redis/go-redis/v9"
)

// RedisClient 實現 Redis 數據庫接口
type RedisClient struct {
	client *redis.Client
	config config.RedisConfig
}

// NewRedisClient 創建新的 Redis 客戶端
func NewRedisClient(cfg config.RedisConfig) *RedisClient {
	return &RedisClient{
		config: cfg,
	}
}

// Connect 建立 Redis 連接
func (c *RedisClient) Connect() error {
	// 創建 Redis 客戶端
	c.client = redis.NewClient(&redis.Options{
		Addr:     c.config.GetRedisAddr(),
		Password: c.config.Password,
		DB:       c.config.DB,
	})

	// 測試連接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := c.client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to connect to Redis: %v", err)
	}

	log.Println("Redis connection established successfully")
	return nil
}

// Close 關閉 Redis 連接
func (c *RedisClient) Close() error {
	if c.client != nil {
		return c.client.Close()
	}
	return nil
}

// HealthCheck 檢查 Redis 健康狀態
func (c *RedisClient) HealthCheck(ctx context.Context) error {
	if c.client == nil {
		return fmt.Errorf("Redis connection is not established")
	}
	return c.client.Ping(ctx).Err()
}

// GetClient 獲取 Redis 客戶端
func (c *RedisClient) GetClient() interface{} {
	return c.client
}
