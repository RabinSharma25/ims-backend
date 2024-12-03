package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabinsharma25/ims-backend/src/controllers"
)

func SetupProductRoutes(router *gin.Engine) {
	controller := controllers.ProductController{}
	user := router.Group("/api/product")
	{
		user.GET("/getProducts", controller.GetProducts)

	}
}
