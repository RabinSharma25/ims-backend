package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabinsharma25/ims-backend/src/controllers"
)

func SetupUserRoutes(router *gin.Engine) {
	controller := controllers.DashboardController{}
	user := router.Group("/api/dashboard")
	{
		// user.POST("/addUserDetails", controller.AddUserDetails)
		// user.PUT("/updateUserDetails", controller.UpdateUserDetails)
		// user.GET("/getUserDetails/:id", controller.GetUserDetails)
		// user.DELETE("/deleteUserDetails/:id", controller.DeleteUserDetails)
		// user.PATCH("/updateUserEmail", controller.UpdateUserEmail)
		user.GET("/getDashboardMetrices", controller.GetDashboardMetrices)

	}
}
