# Glossika Assignment

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
├── migrations/         # Database migrations
├── docs/
│   ├── API.md                               # API 詳細說明文件
│   ├── architecture.plantuml                # 系統架構圖
│   └── Glossika Assignment.postman_collection.json  # Postman API 集合
```

## Getting Started

1. 請先安裝 Docker 與 Docker Compose。
2. 準備 `.env` 檔案（附加在郵件中），放在專案根目錄下。
3. 啟動專案：
   ```bash
   docker-compose --env-file .env up -d
   ```
4. API 測試：透過 `docs/Glossika Assignment.postman_collection.json` 匯入 Postman，提供所有 API 的測試範例