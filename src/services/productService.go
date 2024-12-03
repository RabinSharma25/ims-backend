package services

import (
	"github.com/rabinsharma25/ims-backend/src/dto"
	"github.com/rabinsharma25/ims-backend/src/repositories"
	"github.com/sirupsen/logrus"
)

type ProductService struct{}

var productRepository repositories.ProductRepository = repositories.ProductRepository{}

func (service ProductService) GetProducts() (*dto.GenericResponseDto, error) {
	logrus.Info("UserService.GetProducts")

	res, err := productRepository.GetProducts()

	if err != nil {
		logrus.Error("Failed to get products ")
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
		Message: "Successfully fetched products",
		Data:    res.Data,
	}, nil

}
