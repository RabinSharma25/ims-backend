package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rabinsharma25/ims-backend/src/services"
	"github.com/sirupsen/logrus"
)

type ProductController struct{}

var productService services.ProductService = services.ProductService{}

func (controller ProductController) GetProducts(ctx *gin.Context) {
	logrus.Info("UserController.GetProducts")

	res, err := productService.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
