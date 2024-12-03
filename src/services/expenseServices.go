package services

import (
	"github.com/rabinsharma25/ims-backend/src/dto"
	"github.com/rabinsharma25/ims-backend/src/repositories"
	"github.com/sirupsen/logrus"
)

type ExpenseService struct{}

var expenseRepository repositories.ExpenseRepository = repositories.ExpenseRepository{}

func (service ExpenseService) GetExpensesByCategory() (*dto.GenericResponseDto, error) {
	logrus.Info("ExpenseService.GetExpensesByCategory")

	res, err := expenseRepository.GetExpensesByCategory()

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
		Message: "Successfully expenses",
		Data:    res.Data,
	}, nil

}
