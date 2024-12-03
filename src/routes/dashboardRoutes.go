package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabinsharma25/ims-backend/src/controllers"
)

func SetupDashboardRoutes(router *gin.Engine) {
	controller := controllers.DashboardController{}
	user := router.Group("/api/dashboard")
	{
		user.GET("/getDashboardMetrices", controller.GetDashboardMetrices)

	}
}
