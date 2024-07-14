package repository

import (
	"store-be-golang/models"
	"store-be-golang/structs"
)

func CreateNewOrder(dataOrder structs.OrderInput) (error) {
	order := models.Orders{
		UserID: dataOrder.UserID,
		Address: dataOrder.Address,
		ProductID: dataOrder.ProductID,
	}

	if err := models.DB.Create(&order).Error; err != nil {
		return err
	}

	return nil
}

func UpdatePaymentStatus(dataUpdate structs.UpdatePaymentInput) (error) {
	order := models.Orders{
		PaymentStatus: dataUpdate.PaymentStatus,
	}

	if err := models.DB.Where("order_id = ?", dataUpdate.OrderID).Error; err != nil {
		return err
	}

	if err := models.DB.Model(&order).Updates(dataUpdate).Error; err != nil {
		return err 
	}

	return nil
}