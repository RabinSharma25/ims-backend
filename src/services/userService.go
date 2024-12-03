package services

import (
	"github.com/rabinsharma25/ims-backend/src/dto"
	"github.com/rabinsharma25/ims-backend/src/repositories"
	"github.com/sirupsen/logrus"
)

type UserService struct{}

var userRepository repositories.UserRepository = repositories.UserRepository{}

func (service UserService) GetUsers() (*dto.GenericResponseDto, error) {
	logrus.Info("UserService.GetProducts")

	res, err := userRepository.GetUsers()

	if err != nil {
		logrus.Error("Failed to get users ")
		return &dto.GenericResponseDto{
			Code:    500,
			Success: false,
			Status:  "Failed",
			Message: err.Error(),
			Data:    nil,
		}, err
	}

	return &dto.GenericResponseDto{
		Code:    200,
		Success: true,
		Status:  "Completed",
		Message: "Successfully fetched users",
		Data:    res.Data,
	}, nil

}
