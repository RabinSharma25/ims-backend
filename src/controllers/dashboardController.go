package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rabinsharma25/ims-backend/src/services"
	"github.com/sirupsen/logrus"
)

type DashboardController struct{}

var dashboardService services.DashboardService = services.DashboardService{}

func (controller DashboardController) GetDashboardMetrices(ctx *gin.Context) {
	logrus.Info("UserController.GetDashboardMetrices")

	res, err := dashboardService.GetDashboardMetrices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
