package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabinsharma25/ims-backend/src/controllers"
)

func SetupExpenseRoutes(router *gin.Engine) {
	controller := controllers.ExpenseController{}
	expense := router.Group("/api/expense")
	{
		expense.GET("/getExpensesByCategory", controller.GetExpensesByCategory)

	}
}
