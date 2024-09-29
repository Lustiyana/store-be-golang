package models

// PaymentStatus = PENDING, PAID, FAILED, CONFIRMED

type Orders struct {
	OrderID uint `json:"order_id" gorm:"primaryKey"`
	UserID uint `json:"user_id"`
	Address string `json:"address"`
	CartID uint `json:"card_id"`
	PaymentStatus string `json:"payment_status" gorm:"default:'PENDING'"` 
	ShipingSTatus uint `json:"shiping_status" gorm:"default:0"`
	FinishedStatus uint `json:"finished_status" gorm:"default:0"`
	CreatedAt uint `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt uint `json:"updated_at" gorm:"autoUpdateTime"`
}