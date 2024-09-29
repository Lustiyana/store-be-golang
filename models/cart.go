package models

type Cart struct {
	CartItemID uint `json:"cart_item_id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Quantity uint `json:"quantity" gorm:"default:1`
}