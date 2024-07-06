package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"os"
	"github.com/gin-gonic/gin"
	"store-be-golang/helpers"
	"store-be-golang/structs"
	"store-be-golang/repository"
	"github.com/google/uuid"
)

func CreateNewProduct(c *gin.Context) {
	var input structs.ProductInput

	tokenWithBearer := c.GetHeader("Authorization")

	token, err := helpers.ExtractToken(tokenWithBearer)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}

	data, err := helpers.VerifyToken(token)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}
	
	input.UserID = data.ID

	categoryIDStr := c.PostForm("category_id")
	description := c.PostForm("description")
	priceStr := c.PostForm("price")

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 64)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}

	price, err := strconv.ParseUint(priceStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price"})
		return
	}

	input.CategoryID = uint(categoryID)
	input.Description = description
	input.Price = uint(price)

	productID, err := repository.CreateNewProduct(input)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}

	// Add images
	form, _ := c.MultipartForm()
	files := form.File["images"]

	for _, file := range files {
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
			ProductID: productID,
			Alt: imageURL,
			Url: imageURL,
		}

		err := repository.UploadNewImage(dataImage)
		if err != nil {
			helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
			return
		}
	}

	helpers.GeneralResponse(c, http.StatusOK,true, "Berhasil membuat product baru", nil, nil)
}

func GetAllProduct(c *gin.Context) {
	products, err := repository.GetAllProduct()
	if err != nil {
		panic(err)
		helpers.GeneralResponse(c, http.StatusBadRequest,false, err.Error(), nil, nil)
	}

	helpers.GeneralResponse(c, http.StatusOK,true, "Data berhasil didapatkan", products, nil)
}


func EditProduct(c *gin.Context) {
	var input structs.ProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
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

	id, _ := strconv.Atoi(c.Param("id"))

	product, err := repository.FindProductByID(id)

	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	if product.UserID != dataUser.ID {
		helpers.GeneralResponse(c, http.StatusBadRequest, false, "Invalid Token", nil, nil)
		return
	}

	err = repository.EditProduct(id, input)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	helpers.GeneralResponse(c, http.StatusOK, false, "Berhasil mengedit product", nil, nil)
}

func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

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

	product, err := repository.FindProductByID(id)
	if dataUser.ID != product.UserID {
		helpers.GeneralResponse(c, http.StatusBadRequest, false, "Invalid Token", nil, nil)
		return
	}

	_, err = repository.DeleteProduct(id)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	imagesDeleted, err := repository.DeleteImageByProductID(id)
	if err != nil {
		helpers.GeneralResponse(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	for _, image := range imagesDeleted {
		err = os.Remove("." + image.Url)
		if err != nil {
			helpers.GeneralResponse(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		}
	}

	helpers.GeneralResponse(c, http.StatusOK, true, "Product berhasil dihapus", nil, nil)
}