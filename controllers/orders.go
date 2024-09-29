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

	if err := c.ShouldBindJSON(&input); err != nil {
		panic(err)
		helpers.GeneralResponse(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

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

func GetAllOrders(c *gin.Context) {
	tokenWithBearer := c.GetHeader("Authorization")

	token, err := helpers.ExtractToken(tokenWithBearer)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}

	dataUser, err := helpers.VerifyToken(token)

	orders, err := repository.GetAllOrders(dataUser.ID)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}

	helpers.GeneralResponse(c, http.StatusOK,true, "OK", orders, nil)
}