package models

type Categories struct {
	CategoryID uint `json:"category_id" gorm:"primaryKey"`
	CategoryName string `json:"category_name"`
	Icon string `json:"icon"`
	Products []Products `json:"products" gorm:"foreignKey:CategoryID"`
}