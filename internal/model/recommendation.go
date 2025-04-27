package model

import "time"

// Recommendation 定義推薦項目
type Recommendation struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"size:255;not null"`
	Description string    `json:"description" gorm:"type:text"`
	Score       float64   `json:"score" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
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
