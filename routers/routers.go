package routers

import (
	"github.com/gin-gonic/gin"
	"store-be-golang/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.Static("/public", "./public")

	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)
	router.POST("/products", controllers.CreateNewProduct)
	router.GET("/products", controllers.GetAllProduct)

	return router
}