package services

import (
	"github.com/rabinsharma25/ims-backend/src/dto"
	"github.com/rabinsharma25/ims-backend/src/repositories"
	"github.com/sirupsen/logrus"
)

type DashboardService struct{}

var dashboardRepository repositories.DashboardRepository = repositories.DashboardRepository{}

func (service DashboardService) GetDashboardMetrices() (*dto.GenericResponseDto, error) {
	logrus.Info("UserService.GetDashboardMetrices")

	res, err := dashboardRepository.GetDashboardMetrices()

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
