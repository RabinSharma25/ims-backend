package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rabinsharma25/ims-backend/src/dto"
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

func (controller ProductController) AddProduct(ctx *gin.Context) {
	logrus.Info("ProductController.AddProduct")
	var reqPayload dto.AddProductReqDto
	if err := ctx.ShouldBind(&reqPayload); err != nil {
		logrus.Panic("Error binding the payload")
		ctx.JSON(http.StatusBadRequest, dto.GenericResponseDto{
			Code:    400,
			Success: false,
			Message: "Invalid payload",
			Data:    nil,
		})

		return
	}
	res, err := productService.AddProduct(reqPayload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
