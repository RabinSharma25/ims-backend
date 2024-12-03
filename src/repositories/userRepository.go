package repositories

import (
	"github.com/rabinsharma25/ims-backend/src/db"
	"github.com/rabinsharma25/ims-backend/src/dto"
	"github.com/rabinsharma25/ims-backend/src/models"
	"github.com/sirupsen/logrus"
)

type UserRepository struct{}

func (repo UserRepository) GetUsers() (*dto.GenericResponseDto, error) {
	logrus.Info("UserRepository.GetUsers")
	dbConn := db.GetDatabaseInstance()

	// Slice to store the product list
	UserList := []dto.UserDto{}

	result := dbConn.Model(&models.Users{}).
		Order("name DESC").
		Find(&UserList)

	if result.Error != nil {
		logrus.Error("Error fetching products: ", result.Error)
		return nil, result.Error
	}

	// Return the response
	return &dto.GenericResponseDto{
		Data: map[string]interface{}{

			"UserList": UserList,
		},
	}, nil
}
