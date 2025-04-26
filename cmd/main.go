package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 設置 Gin 模式
	gin.SetMode(gin.ReleaseMode)

	// 創建 Gin 路由
	r := gin.Default()

	// 添加中間件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 健康檢查路由
	r.GET("/health", func(c *gin.Context) {
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

	// 啟動服務器
	log.Println("Server starting on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
