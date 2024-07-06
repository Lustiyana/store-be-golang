package repository

import (
	"store-be-golang/models"
	"store-be-golang/structs"
)

func UploadNewImage(dataImage structs.ImageInput) error {
	image := models.Images{
		ProductID: dataImage.ProductID,
		Alt: dataImage.Alt,
		Url: dataImage.Url,
	}

	if err := models.DB.Create(&image).Error; err != nil {
		return err
	}

	return nil
}

func DeleteImage(id int) (models.Images, error) {
	var image models.Images

	if err := models.DB.Where("image_id = ?", id).First(&image).Error; err != nil {
		return image, err
	}

	models.DB.Delete(&image)

	return image, nil
}

func GetImageByID(id int) (models.Images, error) {
	var image models.Images

	if err := models.DB.Where("image_id = ?", id).First(&image).Error; err != nil {
		return image, err
	}

	return image, nil
}

func DeleteImageByProductID(productID int) ([]models.Images, error) {
	var images []models.Images

	if err := models.DB.Where("product_id = ?", productID).Find(&images).Error; err != nil {
		return images, err
	}

	models.DB.Delete(&images)

	return images, nil
}