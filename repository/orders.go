package repository

import (
	"store-be-golang/models"
	"store-be-golang/structs"
)

func CreateNewOrder(dataOrder structs.OrderInput) (error) {
	order := models.Orders{
		UserID: dataOrder.UserID,
		Address: dataOrder.Address,
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

func GetOrderByPaymentStatus(paymentStatus string) ([]models.Orders, error) {
	var orders []models.Orders

	if err := models.DB.Where("payment_status = ?", paymentStatus).Find(&orders).Error; err != nil {
		return orders, err
	}

	return orders, nil
}

func GetOrderByShipingStatus(shipingStatus string) ([]models.Orders, error) {
	var orders []models.Orders

	if err := models.DB.Where("shiping_status = ?", shipingStatus).Find(&orders).Error; err != nil {
		return orders, err
	}

	return orders, nil
}

func GetOrderByFinishedStatus(finishedStatus string) ([]models.Orders, error) {
	var orders []models.Orders

	if err := models.DB.Where("finished_status = ?", finishedStatus).Find(&orders).Error; err != nil {
		return orders, err
	}

	return orders, nil
}

func GetAllOrders(userID uint) ([]models.Orders, error) {
	var orders []models.Orders
	err := models.DB.
		Joins("JOIN products ON products.product_id = orders.product_id").
		Where("products.user_id = ?", userID).
		Find(&orders).Error
	if err != nil {
		return orders, err
	}

	return orders, nil
}
