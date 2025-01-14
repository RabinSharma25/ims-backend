package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	SetupDashboardRoutes(router)
	SetupProductRoutes(router)
	SetupUserRoutes(router)
	SetupExpenseRoutes(router)
}
