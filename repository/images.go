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