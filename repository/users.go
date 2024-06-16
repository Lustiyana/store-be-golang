package repository

import (
	"gorm.io/gorm"
	"store-be-golang/models"
)

func FindUserByEmail(db *gorm.DB, email string) (models.Users, error) {
	var user models.Users

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}