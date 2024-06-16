package models

type Users struct {
	UserID uint `json:"user_id" gorm:"primary_key"`
	FullName string `json:"fullname"`
	Email string `json:"email"`
	Password string `json:"password"`
	Address string `json:"address"`
	Province string `json:"province"`
	City string `json:"city"`
	PostalCode string `json:"postal_code"`
	Country string `json:"country"`
	PhoneNumber string `json:"phone_number"`
}