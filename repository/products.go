package repository

import (
	"store-be-golang/models"
	"store-be-golang/structs"
	"errors"
	"gorm.io/gorm"
)

func CreateNewProduct(dataProduct structs.ProductInput) (uint, error) {
	product := models.Products{
		CategoryID: dataProduct.CategoryID,
		Price: dataProduct.Price,
		Description: dataProduct.Description,
		UserID: dataProduct.UserID,
		ProductName: dataProduct.ProductName,
	}

	if err := models.DB.Create(&product).Error; err != nil {
		return 0, err
	}

	return product.ProductID, nil
}

func GetAllProduct() ([]models.Products, error) {
	var products []models.Products
  	if err := models.DB.Preload("Images").Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func GetProductsByUserID(userID uint) ([]models.Products, error) {
	var products []models.Products

	if err := models.DB.Preload("Images").Where("user_id = ?", userID).Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func FindProductByID(id int) (models.Products, error) {
	var product models.Products

	if err := models.DB.Preload("Images").Where("product_id = ?", id).First(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func CheckSellerProduct(productID int, userID int) (bool) {
	var product models.Products
	
	err := models.DB.Where("product_id = ? AND user_id = ?", productID, userID).First(&product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}

func EditProduct(id int, dataProduct structs.ProductInput) (error) {
	var product models.Products

	if err := models.DB.Where("product_id = ?", id).First(&product).Error; err != nil {
		return err
	}

	if err := models.DB.Model(&product).Updates(dataProduct).Error; err != nil {
		return err
	}

	return nil
}

func DeleteProduct(id int) (models.Products, error) {
	var product models.Products

	if err := models.DB.Where("product_id = ?", id).First(&product).Error; err != nil {
		return product, err
	}

	models.DB.Delete(&product)

	return product, nil
}