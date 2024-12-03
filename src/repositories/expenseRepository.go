package repositories

import (
	"github.com/rabinsharma25/ims-backend/src/db"
	"github.com/rabinsharma25/ims-backend/src/dto"
	"github.com/rabinsharma25/ims-backend/src/models"
	"github.com/sirupsen/logrus"
)

type ExpenseRepository struct{}

func (repo ExpenseRepository) GetExpensesByCategory() (*dto.GenericResponseDto, error) {
	logrus.Info("ExpenseRepository.GetExpensesByCategory")
	dbConn := db.GetDatabaseInstance()

	// Slice to store the product list
	ExpenseList := []dto.ExpenseByCategoryDto{}

	result := dbConn.Model(&models.ExpenseByCategory{}).
		Order("amount DESC").
		Find(&ExpenseList)

	if result.Error != nil {
		logrus.Error("Error fetching expenses: ", result.Error)
		return nil, result.Error
	}

	// Return the response
	return &dto.GenericResponseDto{
		Data: map[string]interface{}{

			"ExpenseList": ExpenseList,
		},
	}, nil
}
