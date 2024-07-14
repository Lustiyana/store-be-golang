package models

type Categories struct {
	UserID uint `json:"category_id" gorm:"primary_key"`
	CategoryName string `json:"category_name"`
	Icon string `json:"icon"`
	products []Products `json:"products" gorm:"foreignKey:CategoryID"`
}