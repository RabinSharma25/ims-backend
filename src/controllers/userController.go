package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rabinsharma25/ims-backend/src/services"

	"github.com/sirupsen/logrus"
)

type UserController struct{}

var userService services.UserService = services.UserService{}

func (controller UserController) GetUsers(ctx *gin.Context){

	logrus.Info("UserController.GetUsers")
	res,err:= userService.GetUsers()
	
	if(err!=nil){
		ctx.JSON(http.StatusInternalServerError,res)
		return
	}

	ctx.JSON(http.StatusOK,res)

}