package models

type Images struct {
	ImageID uint `json:"image_id" gorm:"primary_key"`
	ProductID uint `json:"product_id"`
	Alt string `json:"alt"`
	Url string `json:"url"`
}