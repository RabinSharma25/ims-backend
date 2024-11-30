package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rabinsharma25/ims-backend/src/dto"
	"github.com/rabinsharma25/ims-backend/src/services"
	"github.com/sirupsen/logrus"
)

type UserController struct{}

var userService services.UserService = services.UserService{}

func (controller UserController) AddUserDetails(ctx *gin.Context) {
	logrus.Info("UserController.AddUserDetails")
	var reqPayload dto.AddUserDetailsReqDto
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
	res, err := userService.AddUserDetails(reqPayload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller UserController) UpdateUserDetails(ctx *gin.Context) {
	logrus.Info("UserController.UpdateUserDetails")
	var reqPayload dto.UpdateUserDetailsReqDto

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
	res, err := userService.UpdateUserDetails(reqPayload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
func (controller UserController) UpdateUserEmail(ctx *gin.Context) {
	logrus.Info("UserController.UpdateUserEmail")
	var reqPayload dto.UpdateUserEmailReqDto

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
	res, err := userService.UpdateUserEmail(reqPayload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller UserController) GetUserDetails(ctx *gin.Context) {
	logrus.Info("UserController.GetUserDetails")

	// Get the ID from the URL and convert it to an integer
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.GenericResponseDto{
			Code:    400,
			Success: false,
			Message: "Invalid ID parameter",
			Data:    nil,
		})
		return
	}
	res, err := userService.GetUserDetails(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller UserController) DeleteUserDetails(ctx *gin.Context) {
	logrus.Info("UserController.DeleteUserDetails")

	// Get the ID from the URL and convert it to an integer
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.GenericResponseDto{
			Code:    400,
			Success: false,
			Message: "Invalid ID parameter",
			Data:    nil,
		})
		return
	}
	res, err := userService.DeleteUserDetails(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
