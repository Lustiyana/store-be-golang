package models

type Products struct {
	ProductID uint `json:"product_id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`
	CategoryID uint `json:"category_id"`
	ProductName string `json:"product_name"`
	Price uint `json:"price"`
	Description string `json:"description"`
	Images []Images `json:"images" gorm:"foreignKey:ProductID"`
}