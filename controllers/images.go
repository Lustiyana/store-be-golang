package controllers

import (
	"strings"
	"fmt"
	"net/http"
	"strconv"
	"os"
	"store-be-golang/structs"
	"store-be-golang/repository"
	"store-be-golang/helpers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadNewImage(c *gin.Context) {
	file, _ := c.FormFile("image")
	productID, _ := strconv.ParseUint(c.PostForm("product_id"), 10, 64)

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

	product, err := repository.FindProductByID(int(productID))
	if dataUser.ID != product.UserID {
		helpers.GeneralResponse(c, http.StatusBadRequest, false, "Invalid Token", nil, nil)
		return
	}

	uniqueID := uuid.New()
	filename := strings.Replace(uniqueID.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	imageName := fmt.Sprintf("%s.%s", filename, fileExt)

	err = c.SaveUploadedFile(file, fmt.Sprintf("./public/images/%s", imageName))
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}

	imageURL := fmt.Sprintf("/public/images/%s", imageName)

	dataImage := structs.ImageInput{
		ProductID: uint(productID),
		Alt: imageURL,
		Url: imageURL,
	}

	err = repository.UploadNewImage(dataImage)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	helpers.GeneralResponse(c, http.StatusOK, true, "Gambar berhasil ditambahkan", nil, nil)
}

func DeleteImage(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	dataImage, err := repository.GetImageByID(id)
	if err != nil {
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

	product, err := repository.FindProductByID(int(dataImage.ProductID))
	if dataUser.ID != product.UserID {
		helpers.GeneralResponse(c, http.StatusBadRequest, false, "Invalid Token", nil, nil)
		return
	}

	imageDeleted, err := repository.DeleteImage(id)

	err = os.Remove("." + imageDeleted.Url)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	helpers.GeneralResponse(c, http.StatusOK, true, "Gambar berhasil dihapus", nil, nil)
}