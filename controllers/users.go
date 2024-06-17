package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"store-be-golang/repository"
	"gorm.io/gorm"
	"store-be-golang/structs"
	"store-be-golang/helpers"
)

func Register(c *gin.Context) {
	var dataUser structs.RegisterInput

	if err := c.ShouldBindJSON(&dataUser); err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}

	_, err := repository.FindUserByEmail(dataUser.Email)
	if err == nil {
		helpers.GeneralResponse(c, http.StatusBadRequest, false, "Email sudah terdaftar", nil, nil)
		return
	}

	if dataUser.Password != dataUser.ConfirmPassword {
		helpers.GeneralResponse(c, http.StatusBadRequest, false, "Password tidak sama", nil, nil)
		return
	}

	hashPassword, err := helpers.GenerateHashPassword(dataUser.Password)

	dataUser.Password = hashPassword

	if err := repository.CreateUser(dataUser); err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	helpers.GeneralResponse(c, http.StatusOK,true, "Akun berhasil dibuat", nil, nil)
}

func Login(c *gin.Context) {
	var dataUser structs.LoginInput

	if err := c.ShouldBindJSON(&dataUser); err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}

	res, err := repository.FindUserByEmail(dataUser.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.GeneralResponse(c, http.StatusBadRequest,false, "Email tidak ditemukan", nil, nil)
			return
		}

		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}

	err = helpers.ValidatePassword(res.Password, dataUser.Password)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest,false, "Password salah", nil, nil)
		return
	}

	token, err := helpers.GenerateToken(res.UserID, dataUser.Email, dataUser.Password)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}

	tokenResult := map[string]interface{}{
		"token": token,
	}

	helpers.GeneralResponse(c, http.StatusOK,true, "Berhasil Masuk", tokenResult, nil)
}