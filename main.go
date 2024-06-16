package main

import (
	"github.com/gin-gonic/gin"
	"store-be-golang/controllers"
	"store-be-golang/models"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func main() {
	r := gin.Default()

	r.POST("/login", controllers.Login)

	models.ConnectDatabase()
	
	hash, err := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.MinCost)
    if err != nil {
        fmt.Println(err)
    }

		fmt.Println(string(hash))
	r.Run()
}