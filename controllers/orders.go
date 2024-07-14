package controllers

import(
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"store-be-golang/repository"
	"store-be-golang/structs"
	"store-be-golang/helpers"
)

func CreateNewOrder(c *gin.Context) {
	var input structs.OrderInput

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
	
	address := c.PostForm("address")

	productID, _ := strconv.ParseUint(c.PostForm("product_id"), 10, 64)

	input.Address = address
	input.ProductID = uint(productID)
	input.UserID = dataUser.ID

	err = repository.CreateNewOrder(input)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}

	helpers.GeneralResponse(c, http.StatusOK,true, "Berhasil membuat pesanan", nil, nil)
}

func UpdatePaymentStatus(c *gin.Context) {
	var input structs.UpdatePaymentInput

	orderID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	paymentStatus := c.PostForm("payment_status")

	input.OrderID = uint(orderID)
	input.PaymentStatus = paymentStatus

	err := repository.UpdatePaymentStatus(input)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}

	helpers.GeneralResponse(c, http.StatusOK,true, "Payment berhasil diupdate", nil, nil)
}
