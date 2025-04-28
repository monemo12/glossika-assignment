# Glossika Assignment

This is a Go project for the recommendation system.

## Project Layout

```
glossika-assignment/
├── cmd/                # 主程式
├── configs/            # 設定檔
├── internal/
│   ├── config/         # 設定結構與讀取
│   ├── database/       # 資料庫連線、操作與種子資料（MySQL、Redis）
│   ├── handler/        # HTTP 路由與請求處理
│   ├── middleware/     # Gin 中介層（如 JWT 驗證）
│   ├── model/          # 資料結構定義（User、Recommendation）
│   ├── repository/     # 資料存取層，與資料庫互動
│   ├── service/        # 商業邏輯（用戶、推薦、Email 服務）
│   └── utils/          # 通用工具（密碼、Email、JWT 等）
├── migrations/
├── docs/
```

## Getting Started

1. 請先安裝 Docker 與 Docker Compose。
2. Clone 此 repository。
3. 準備 `.env` 檔案（檔案內容會由我以郵件提供），放在專案根目錄下。
4. 啟動專案：
   ```bash
   docker-compose --env-file .env up -d
   ```

## System Architecture