package controllers

import (
	"net/http"
	"strconv"
	"store-be-golang/repository"
	"store-be-golang/helpers"
	"github.com/gin-gonic/gin"
)

func GetCategoryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	dataCategories, err := repository.GetCategoryByID(id)
	if err != nil {
		// panic(err)
		helpers.GeneralResponse(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	helpers.GeneralResponse(c, http.StatusOK, true, "Berhasil mendapatkan kategori", dataCategories, nil)
}

func GetAllCategories(c *gin.Context) {
	categories, err := repository.GetAllCategories()
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	helpers.GeneralResponse(c, http.StatusOK, true, "Data berhasil didapatkan", categories, nil)
}