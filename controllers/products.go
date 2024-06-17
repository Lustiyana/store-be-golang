package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
	"store-be-golang/helpers"
	"store-be-golang/structs"
	"store-be-golang/repository"
	"github.com/google/uuid"
)

func CreateNewProduct(c *gin.Context) {
	var input structs.ProductInput

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

	fmt.Println(productID)

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