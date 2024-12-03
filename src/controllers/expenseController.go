package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rabinsharma25/ims-backend/src/services"

	"github.com/sirupsen/logrus"
)

type ExpenseController struct{}

var expenseService services.ExpenseService = services.ExpenseService{}

func (controller ExpenseController) GetExpensesByCategory(ctx *gin.Context){

	logrus.Info("ExpenseController.GetExpensesByCategory")
	res,err:= expenseService.GetExpensesByCategory()
	
	if(err!=nil){
		ctx.JSON(http.StatusInternalServerError,res)
		return
	}

	ctx.JSON(http.StatusOK,res)

}