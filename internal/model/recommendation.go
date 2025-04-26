package model

// RecommendationItem 定義推薦項目
type RecommendationItem struct {
	ID          string
	Title       string
	Description string
}

// RecommendationRequest 定義推薦請求
type RecommendationRequest struct {
	UserID string
	Limit  int
	Offset int
}

// RecommendationResponse 定義推薦響應
type RecommendationResponse struct {
	Items    []*RecommendationItem
	Total    int
	NextPage bool
}
