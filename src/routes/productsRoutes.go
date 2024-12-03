package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabinsharma25/ims-backend/src/controllers"
)

func SetupProductRoutes(router *gin.Engine) {
	controller := controllers.ProductController{}
	product := router.Group("/api/product")
	{
		product.GET("/getProducts", controller.GetProducts)
		product.POST("/addProduct", controller.AddProduct)

	}
}
