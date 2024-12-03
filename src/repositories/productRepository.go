package repositories

import (
	"github.com/rabinsharma25/ims-backend/src/db"
	"github.com/rabinsharma25/ims-backend/src/dto"
	"github.com/rabinsharma25/ims-backend/src/models"
	"github.com/sirupsen/logrus"
)

type ProductRepository struct{}

func (repo ProductRepository) GetProducts() (*dto.GenericResponseDto, error) {
	logrus.Info("ProductRepository.GetProducts")
	dbConn := db.GetDatabaseInstance()

	// Slice to store the product list
	ProductList := []dto.ProductDto{}

	// Fetch the first 15 records sorted in descending order by stock_quantity
	result := dbConn.Model(&models.Products{}).
		Order("stock_quantity DESC"). 
		Find(&ProductList)

	if result.Error != nil {
		logrus.Error("Error fetching products: ", result.Error)
		return nil, result.Error
	}

	// Return the response
	return &dto.GenericResponseDto{
		Data: []map[string]interface{}{
			{
				"ProductList": ProductList,
			},
		},
	}, nil
}
