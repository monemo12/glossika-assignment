package database

import (
	"context"

	"gorm.io/gorm"
)

// Database 定義數據庫通用接口
type IDatabase interface {
	// Connect 建立數據庫連接
	Connect() error
	// Close 關閉數據庫連接
	Close() error
	// HealthCheck 檢查數據庫健康狀態
	HealthCheck(ctx context.Context) error
}

// MySQLDatabase 定義 MySQL 數據庫接口
type MySQLDatabase interface {
	IDatabase
	// GetDB 獲取 GORM 實例
	GetDB() *gorm.DB
}

// RedisDatabase 定義 Redis 數據庫接口
type RedisDatabase interface {
	IDatabase
	// GetClient 獲取 Redis 客戶端
	GetClient() interface{}
}
