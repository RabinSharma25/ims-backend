package repositories

import (
	"github.com/rabinsharma25/ims-backend/src/db"
	"github.com/rabinsharma25/ims-backend/src/dto"
	"github.com/rabinsharma25/ims-backend/src/models"
	"github.com/sirupsen/logrus"
)

type DashboardRepository struct{}

func (repo DashboardRepository) GetDashboardMetrices() (*dto.GenericResponseDto, error) {
	logrus.Info("UserRepository.GetDashboardMetrices")
	dbConn := db.GetDatabaseInstance()

	// Slice to store the product list
	ProductList := []dto.ProductDto{}
	SalesSummaryList := []dto.SalesSummaryDto{}
	PurchaseSummaryList := []dto.PurchaseSummaryDto{}
	ExpenseSummaryList := []dto.ExpenseSummaryDto{}
	ExpenseByCategoryList := []dto.ExpenseByCategoryDto{}

	// Fetch the first 15 records sorted in descending order by stock_quantity
	result := dbConn.Model(&models.Products{}).
		Order("stock_quantity DESC"). // Sorting by stock_quantity in descending order
		Limit(15).
		Find(&ProductList)

	if result.Error != nil {
		logrus.Error("Error fetching products: ", result.Error)
		return nil, result.Error
	}

	result = dbConn.Model(&models.SalesSummary{}).
		Order("date DESC"). // Sorting by stock_quantity in descending order
		Limit(5).
		Find(&SalesSummaryList)

	if result.Error != nil {
		logrus.Error("Error fetching products: ", result.Error)
		return nil, result.Error
	}

	result = dbConn.Model(&models.PurchaseSummary{}).
		Order("date DESC"). // Sorting by stock_quantity in descending order
		Limit(5).
		Find(&PurchaseSummaryList)

	if result.Error != nil {
		logrus.Error("Error fetching products: ", result.Error)
		return nil, result.Error
	}

	result = dbConn.Model(&models.ExpenseSummary{}).
		Order("date DESC"). // Sorting by stock_quantity in descending order
		Limit(5).
		Find(&ExpenseSummaryList)

	if result.Error != nil {
		logrus.Error("Error fetching products: ", result.Error)
		return nil, result.Error
	}

	result = dbConn.Model(&models.ExpenseByCategory{}).
		Select(`id, 
		        expense_summary_id, 
		        category, 
		        CAST(amount AS TEXT) AS amount, 
		        date`).
		Order("date DESC").
		Limit(5).
		Find(&ExpenseByCategoryList)

	if result.Error != nil {
		logrus.Error("Error fetching expense by category: ", result.Error)
		return nil, result.Error
	}

	// Return the response
	return &dto.GenericResponseDto{
		Data: []map[string]interface{}{
			{"ProductList": ProductList},
			{ "SalesSummaryList": SalesSummaryList},
			{ "PurchaseSummaryList": PurchaseSummaryList},
			{ "ExpenseSummaryList": ExpenseSummaryList},
			{ "ExpenseByCategoryList": ExpenseByCategoryList},
		},
	}, nil
}
