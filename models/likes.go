package models

type Likes struct {
	LikeID uint `json:"like_id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`
	ProductID string `json:"produt_id"`
}