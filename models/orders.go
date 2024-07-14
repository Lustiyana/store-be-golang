package models

type Orders struct {
	OrderID uint `json:"order_id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`
	Address string `json:"address"`
	ProductID uint `json:"product_id"`
	PaymentStatus string `json:"payment_status" gorm:"default:'PENDING'"`
	ShipingSTatus uint `json:"shiping_status" gorm:"default:1"`
	FinishedStatus uint `json:"finished_status" gorm:"default:1"`
	CreatedAt uint `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt uint `json:"updated_at" gorm:"autoUpdateTime"`
}