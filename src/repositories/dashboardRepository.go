package repositories

import (
	"errors"

	"github.com/rabinsharma25/ims-backend/src/db"
	"github.com/rabinsharma25/ims-backend/src/dto"
	"github.com/rabinsharma25/ims-backend/src/models"
	"github.com/sirupsen/logrus"
)

func (repo UserRepository) GetDashboardMetrices() (*dto.GenericResponseDto, error) {
	logrus.Info("UserRepository.GetDashboardMetrices")
	dbConn := db.GetDatabaseInstance()

	userDetails := models.Users{}

	result := dbConn.Model(&models.Users{}).Where("id = ?", userId).Find(&userDetails)
	if result.Error != nil {
		logrus.Error("Error adding user details")
		return nil, errors.New("error fetching store owner details")
	}

	res := dto.GetUserDetailsResDto{
		Id:        int(userDetails.Id),
		UserName:  userDetails.UserName,
		FirstName: userDetails.FirstName,
		LastName:  userDetails.LastName,
		Email:     userDetails.Email,
	}

	return &dto.GenericResponseDto{
		Data: res,
	}, nil
}
