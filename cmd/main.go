package main

import (
	"fmt"
	"log"

	"glossika-assignment/internal/config"
	"glossika-assignment/internal/database"
	"glossika-assignment/internal/handler"
	"glossika-assignment/internal/repository"
	"glossika-assignment/internal/service"
	"glossika-assignment/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加載配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化 JWT
	utils.InitJWT(cfg.JWT)

	// 初始化 MySQL 連接
	mysqlClient := database.NewMySQLClient(cfg.MySQL)
	if err := mysqlClient.Connect(); err != nil {
		log.Fatalf("Failed to initialize MySQL: %v", err)
	}
	defer func() {
		if err := mysqlClient.Close(); err != nil {
			log.Printf("Error closing MySQL connection: %v", err)
		}
	}()

	// 執行數據庫 seed - 生成100個推薦項目
	if err := database.SeedRecommendations(mysqlClient.GetDB(), 100); err != nil {
		log.Printf("Warning: Failed to seed database: %v", err)
	}

	// 初始化 Redis 連接
	redisClient := database.NewRedisClient(cfg.Redis)
	if err := redisClient.Connect(); err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}
	defer func() {
		if err := redisClient.Close(); err != nil {
			log.Printf("Error closing Redis connection: %v", err)
		}
	}()

	// 設置 Gin 模式
	if cfg.Server.Environment == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 創建 Gin 路由
	r := gin.Default()

	// 添加中間件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 健康檢查路由
	r.GET("/health", func(c *gin.Context) {
		// 檢查數據庫健康狀態
		if err := mysqlClient.HealthCheck(c.Request.Context()); err != nil {
			c.JSON(500, gin.H{
				"status": "error",
				"error":  "MySQL health check failed",
			})
			return
		}

		if err := redisClient.HealthCheck(c.Request.Context()); err != nil {
			c.JSON(500, gin.H{
				"status": "error",
				"error":  "Redis health check failed",
			})
			return
		}

		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// 根路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Glossika Assignment API!",
		})
	})

	// 初始化 API 路由
	apiGroup := r.Group("/api/v1")

	// Initialize repositories
	userRepo := repository.NewUserRepository(mysqlClient)
	recommendationRepo := repository.NewRecommendationRepository(mysqlClient, redisClient)

	// Initialize services
	emailService := service.NewEmailService()
	userService := service.NewUserService(userRepo, emailService)
	recommendationService := service.NewRecommendationService(recommendationRepo)

	// Initialize and setup handlers
	userHandler := handler.NewUserHandler(userService)
	userHandler.SetupRoutes(apiGroup)

	recommendationHandler := handler.NewRecommendationHandler(recommendationService)
	recommendationHandler.SetupRoutes(apiGroup)

	// 啟動服務器
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server starting on %s...", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
