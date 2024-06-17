package repository

import (
	"store-be-golang/models"
	"store-be-golang/structs"
)

func FindUserByEmail(email string) (models.Users, error) {
	var user models.Users

	if err := models.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(dataUser structs.RegisterInput) (error) {
	user := models.Users{
		Email: dataUser.Email,
		Password: dataUser.Password,
		FullName: dataUser.FullName,
		Address: dataUser.Address,
		Province: dataUser.Province,
		City: dataUser.City,
		PostalCode: dataUser.PostalCode,
		Country: dataUser.Country,
		PhoneNumber: dataUser.PhoneNumber,
	}

	if err := models.DB.Create(&user).Error; err != nil {
		return err
	}
	
	return nil
}