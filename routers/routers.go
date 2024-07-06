package routers

import (
	"github.com/gin-gonic/gin"
	"store-be-golang/controllers"
	"store-be-golang/middleware"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	authGroup := router.Group("/")
	authGroup.Use(middleware.AuthMiddleware())

	router.Static("/public", "./public")

	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)

	authGroup.POST("/products", controllers.CreateNewProduct)
	authGroup.GET("/products", controllers.GetAllProduct)
	authGroup.PUT("/products/:id", controllers.EditProduct)
	authGroup.DELETE("/products/:id", controllers.DeleteProduct)

	authGroup.POST("/images", controllers.UploadNewImage)
	authGroup.DELETE("/images/:id", controllers.DeleteImage)

	return router
}