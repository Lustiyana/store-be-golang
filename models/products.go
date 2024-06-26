package models

type Products struct {
	ProductID uint `json:"product_id" gorm:"primary_key"`
	CategoryID uint `json:"category_id"`
	Price uint `json:"price"`
	Description string `json:"description"`
	Images []Images `json:"images" gorm:"foreignKey:ProductID"`
}