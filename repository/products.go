package repository

import (
	"store-be-golang/models"
	"store-be-golang/structs"
)

func CreateNewProduct(dataProduct structs.ProductInput) (uint, error) {
	product := models.Products{
		CategoryID: dataProduct.CategoryID,
		Price: dataProduct.Price,
		Description: dataProduct.Description,
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