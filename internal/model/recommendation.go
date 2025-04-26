package model

// RecommendationItem 定義推薦項目
type RecommendationItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// RecommendationRequest 定義推薦請求
type RecommendationRequest struct {
	Limit  int `form:"limit" json:"limit"`
	Offset int `form:"offset" json:"offset"`
}

// RecommendationResponse 定義推薦響應
type RecommendationResponse struct {
	Items    []*RecommendationItem `json:"items"`
	Total    int                   `json:"total"`
	NextPage bool                  `json:"nextPage"`
}
