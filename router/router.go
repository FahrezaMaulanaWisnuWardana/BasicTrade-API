package router

import (
	"BasicTrade-API/controllers"
	"BasicTrade-API/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	// Auth
	authRouter := router.Group("/auth")
	{
		authRouter.POST("register", controllers.Register)
		authRouter.POST("login", controllers.Login)
	}
	// Products
	productRouter := router.Group("/products")
	{
		productRouter.GET("/", controllers.GetProducts)
		productRouter.GET(":uuid", controllers.ProductByUUID)

		variantRouter := productRouter.Group("/variants")
		{
			variantRouter.GET("/", controllers.GetVariants)
			variantRouter.GET("/:uuid", controllers.VariantsByUUID)

			variantRouter.Use(middleware.Authentication())
			variantRouter.POST("/", controllers.AddVariants)
			variantRouter.PUT("/:uuid", controllers.UpdateVariants)
			variantRouter.DELETE("/:uuid", controllers.DeleteVariants)
		}

		productRouter.Use(middleware.Authentication())
		productRouter.POST("/", controllers.AddProduct)
		productRouter.PUT(":uuid", middleware.ProductAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE(":uuid", middleware.ProductAuthorization(), controllers.DeleteProduct)
	}
	return router
}
