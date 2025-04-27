package model

// Recommendation 定義推薦項目
type Recommendation struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Score       float64 `json:"score"`
}

// RecommendationRequest 定義推薦請求
type RecommendationRequest struct {
	Limit  int `form:"limit" json:"limit"`
	Offset int `form:"offset" json:"offset"`
}

// RecommendationResponse 定義推薦響應
type RecommendationResponse struct {
	Items    []*Recommendation `json:"items"`
	Total    int               `json:"total"`
	NextPage bool              `json:"nextPage"`
}
