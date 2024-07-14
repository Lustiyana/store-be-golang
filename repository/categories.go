package repository

import (
	"store-be-golang/models"
)

func GetCategoryByID(categoryID int) (models.Categories, error) {
	var category models.Categories

	if err := models.DB.Where("category_id = ?", categoryID).Find(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func GetAllCategories() ([]models.Categories, error) {
	var categories []models.Categories
	if err := models.DB.Preload("Products").Find(&categories).Error; err != nil {
		return categories, err
	}

	return categories, nil
}