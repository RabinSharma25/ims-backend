package services

import (
	"github.com/rabinsharma25/ims-backend/src/dto"
	"github.com/sirupsen/logrus"
)

func (service UserService) GetDashboardMetrices() (*dto.GenericResponseDto, error) {
	logrus.Info("UserService.GetDashboardMetrices")

	res, err := userRepository.GetDashboardMetrices()

	if err != nil {
		logrus.Error("Failed to get details ")
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
		Message: "Successfully fetched dashboard metrices",
		Data:    res.Data,
	}, nil

}
