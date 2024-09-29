package controllers

import (
	"net/http"
	"strconv"
	"store-be-golang/repository"
	"store-be-golang/helpers"
	"github.com/gin-gonic/gin"
	"store-be-golang/structs"
)

func InsertProductIntoCart(c *gin.Context) {
	var input structs.CartInput
	productID, _ := strconv.Atoi(c.PostForm("product_id"))
	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	tokenWithBearer := c.GetHeader("Authorization")

	
	token, err := helpers.ExtractToken(tokenWithBearer)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}
	
	dataUser, err := helpers.VerifyToken(token)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}

	input.ProductID = uint(productID)
	input.UserID = uint(dataUser.ID)
	input.Quantity = uint(quantity)
	
	err = repository.InsertProductIntoCart(input)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}
	
	helpers.GeneralResponse(c, http.StatusOK, false, "Data berhasil ditambahkan ke keranjang", nil, nil)
}

func DeleteProductFromCart(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := repository.DeleteProductFromCart(id)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	helpers.GeneralResponse(c, http.StatusOK, false, "Data berhasil dihapus dari keranjang", nil, nil)
}

func UpdateQuantity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	updateType := c.Query("type")
	
	err := repository.UpdateQuantity(id, updateType)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	helpers.GeneralResponse(c, http.StatusOK, false, "Data berhasil diupdate", nil, nil)
}