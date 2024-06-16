package models

type Products struct {
	ProductID uint `json:"product_id" gorm:"primary_key"`
	CategoryID uint `json:"category_id"`
	Description string `json:"description"`
	ThumbnailID string `json:"thumbnail_id"`
}