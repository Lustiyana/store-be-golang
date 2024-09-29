package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "store-be-golang/controllers"
    "store-be-golang/middleware"
)

func StartServer() *gin.Engine {
    router := gin.Default()

    corsConfig := cors.Config{
        AllowOrigins:     []string{"*"},  // Allow all origins
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // Allow specific methods
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }

    router.Use(cors.New(corsConfig))

    authGroup := router.Group("/")
    authGroup.Use(middleware.AuthMiddleware())

    router.Static("/public", "./public")

    router.POST("/login", controllers.Login)
    router.POST("/register", controllers.Register)

    authGroup.POST("/products", controllers.CreateNewProduct)
    authGroup.GET("/products", controllers.GetAllProduct)
    authGroup.PUT("/products/:id", controllers.EditProduct)
    authGroup.DELETE("/products/:id", controllers.DeleteProduct)
    authGroup.GET("/products/:id", controllers.ProductDetail)
    authGroup.GET("/products/me", controllers.GetMyProducts)

    authGroup.POST("/images", controllers.UploadNewImage)
    authGroup.DELETE("/images/:id", controllers.DeleteImage)

    authGroup.GET("/categories/:id", controllers.GetCategoryByID)
    authGroup.GET("/categories", controllers.GetAllCategories)

    authGroup.POST("/orders", controllers.CreateNewOrder)
    authGroup.GET("/orders", controllers.GetAllOrders)
    // Update status order
    // yang bisa update paymentStatus cuma seller
    // yang bisa finish status user

    authGroup.POST("/cart", controllers.InsertProductIntoCart)
    authGroup.DELETE("/cart", controllers.DeleteProductFromCart)
    authGroup.PUT("/cart/:id", controllers.UpdateQuantity)

    // like
    // comment

    return router
}
