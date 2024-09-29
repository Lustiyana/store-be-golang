package models

type OrderedList struct {
	OrderedListID uint `json:"ordered_list_id" gorm:"primaryKey"`
	ProductID string `json:"produt_id"`
	Quantity uint `json:"quantity"`
	OrderID uint `json:OrderID`
}