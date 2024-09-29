package repository

import (
	"store-be-golang/models"
	"store-be-golang/structs"
	"gorm.io/gorm"
)

func InsertProductIntoCart(data structs.CartInput) (error) {
	cart := models.Cart{
		ProductID: data.ProductID,
		Quantity: data.Quantity,
		UserID: data.UserID,
	}

	if err := models.DB.Create(&cart).Error; err != nil {
		return err
	} 

	return nil
}

func DeleteProductFromCart(id int) (error) {
	var cart models.Cart

	if err := models.DB.Delete(&cart).Error; err != nil {
		return err
	}

	return nil
}

func UpdateQuantity(id int, updateType string) (error) {
	var cart models.Cart
	var operator string

	switch updateType {
	case "decrement":
		operator = "-"
	default:
		operator = "+"
	}

	if err := models.DB.Model(&cart).Where("cart_item_id", id).UpdateColumn("quantity", gorm.Expr("quantity ? ?", operator, 1)).Error; err != nil {
		return err
	}

	return nil
}