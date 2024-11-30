package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (controller UserController) GetDashboardMetrices(ctx *gin.Context) {
	logrus.Info("UserController.GetDashboardMetrices")

	res, err := userService.GetDashboardMetrices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
