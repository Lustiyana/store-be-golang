package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"store-be-golang/repository"
	"store-be-golang/models"
	"store-be-golang/helpers"
	"gorm.io/gorm"
	"store-be-golang/structs"
)

// func Register(c *gin.Context) {
// 	var dataUser structs.RegisterInput

// 	if err := c.ShouldBindJSON(&dataUser); err != nil {
// 		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
// 		return
// 	}
// }

func Login(c *gin.Context) {
	var dataUser structs.LoginInput

	if err := c.ShouldBindJSON(&dataUser); err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}

	res, err := repository.FindUserByEmail(models.DB, dataUser.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.GeneralResponse(c, http.StatusBadRequest,false, "Email tidak ditemukan", nil, nil)
			return
		}

		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}

	if res.Password != dataUser.Password {
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