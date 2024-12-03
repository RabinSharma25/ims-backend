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
		Data: map[string]interface{}{

			"ProductList": ProductList,
		},
	}, nil
}


func (repo ProductRepository) AddProduct(payload dto.AddProductReqDto) (*dto.GenericResponseDto, error) {
	logrus.Info("ProductRepository.AddProduct")
	dbConn := db.GetDatabaseInstance()


	newProduct := models.Products{
		Id: payload.Id,
		Name:payload.Name,
		Price: payload.Price,
		Rating: payload.Rating,
		StockQuantity: payload.StockQuantity,
	}

	result := dbConn.Model(&models.Products{}).Create(&newProduct)
	if result.Error != nil {
		logrus.Error("Error adding product to database")
		return nil, result.Error
	}
	return nil, nil
}