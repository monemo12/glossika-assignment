package config

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

// Config 定義完整的配置結構
type Config struct {
	Server ServerConfig `mapstructure:"server"`
	MySQL  MySQLConfig  `mapstructure:"mysql"`
	Redis  RedisConfig  `mapstructure:"redis"`
	JWT    JWTConfig    `mapstructure:"jwt"`
}

// ServerConfig 定義服務器配置
type ServerConfig struct {
	Port        string `mapstructure:"port"`
	Host        string `mapstructure:"host"`
	Environment string `mapstructure:"environment"`
}

// MySQLConfig 定義 MySQL 配置
type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"database"`
}

// RedisConfig 定義 Redis 配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// JWTConfig 定義 JWT 配置
type JWTConfig struct {
	Secret        string `mapstructure:"secret"`
	ExpireMinutes int    `mapstructure:"expireMinutes"`
}

// LoadConfig 加載配置
func LoadConfig() (*Config, error) {
	// 獲取項目根目錄
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(filepath.Dir(b)))

	// 初始化 Viper
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(filepath.Join(projectRoot, "configs"))

	// 設置環境變量前綴
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()

	// 讀取配置文件
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	// 將配置綁定到結構體
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	fmt.Printf("config: %v\n", config)
	return &config, nil
}

// GetMySQLDSN 獲取 MySQL DSN
func (c *MySQLConfig) GetMySQLDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.DBName)
}

// GetRedisAddr 獲取 Redis 地址
func (c *RedisConfig) GetRedisAddr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
