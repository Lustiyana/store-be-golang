package models

type Comments struct {
	CommentID uint `json:"comment_id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`
	ProductID string `json:"produt_id"`
	Comment string `json:"comment"`
}