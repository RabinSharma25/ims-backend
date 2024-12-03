package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabinsharma25/ims-backend/src/controllers"
)

func SetupUserRoutes(router *gin.Engine) {
	controller := controllers.UserController{}
	user := router.Group("/api/user")
	{
		user.GET("/getUsers", controller.GetUsers)

	}
}
