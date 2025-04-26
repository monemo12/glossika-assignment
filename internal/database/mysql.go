package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"glossika-assignment/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// MySQLClient 實現 MySQL 數據庫接口
type MySQLClient struct {
	db     *gorm.DB
	config config.MySQLConfig
}

// NewMySQLClient 創建新的 MySQL 客戶端
func NewMySQLClient(cfg config.MySQLConfig) *MySQLClient {
	return &MySQLClient{
		config: cfg,
	}
}

// Connect 建立 MySQL 連接
func (c *MySQLClient) Connect() error {
	// 配置 GORM
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// 連接 MySQL
	var err error
	c.db, err = gorm.Open(mysql.Open(c.config.GetMySQLDSN()), gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL: %v", err)
	}

	// 設置連接池
	sqlDB, err := c.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get MySQL instance: %v", err)
	}

	// 設置連接池參數
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("MySQL connection established successfully")
	return nil
}

// Close 關閉 MySQL 連接
func (c *MySQLClient) Close() error {
	if c.db != nil {
		sqlDB, err := c.db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

// HealthCheck 檢查 MySQL 健康狀態
func (c *MySQLClient) HealthCheck(ctx context.Context) error {
	if c.db == nil {
		return fmt.Errorf("MySQL connection is not established")
	}

	sqlDB, err := c.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.PingContext(ctx)
}

// GetDB 獲取 GORM 實例
func (c *MySQLClient) GetDB() *gorm.DB {
	return c.db
}
