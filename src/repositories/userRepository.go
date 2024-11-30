package repositories

import (
	"errors"

	"github.com/rabinsharma25/ims-backend/src/db"
	"github.com/rabinsharma25/ims-backend/src/dto"
	"github.com/rabinsharma25/ims-backend/src/models"
	"github.com/sirupsen/logrus"
)

type UserRepository struct{}

func (repo UserRepository) AddUserDetails(payload dto.AddUserDetailsReqDto) (*dto.GenericResponseDto, error) {
	logrus.Info("UserRepository.AddUserDetails")
	dbConn := db.GetDatabaseInstance()

	// Check if email already exists
	var count int64
	dbConn.Model(&models.Users{}).Where("email = ?", payload.Email).Count(&count)
	if count > 0 {
		logrus.Warn("User already exists in the database")
		return nil, errors.New("user already exists")
	}
	customerDetails := models.Users{
		UserName:  payload.UserName,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
	}

	result := dbConn.Model(&models.Users{}).Create(&customerDetails)
	if result.Error != nil {
		logrus.Error("Error inserting newUser to database")
		return nil, result.Error
	}
	return nil, nil
}

func (repo UserRepository) UpdateUserDetails(payload dto.UpdateUserDetailsReqDto) (*dto.GenericResponseDto, error) {
	logrus.Info("UserRepository.UpdateUserDetails")
	dbConn := db.GetDatabaseInstance()

	newUser := models.Users{
		UserName:  payload.UserName,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
	}

	result := dbConn.Model(&models.Users{}).Where("id = ?", payload.Id).Updates(&newUser)
	if result.Error != nil {
		logrus.Error("Error updating the details ")
		return nil, result.Error
	}

	return nil, nil
}

func (repo UserRepository) UpdateUserEmail(payload dto.UpdateUserEmailReqDto) (*dto.GenericResponseDto, error) {
	logrus.Info("UserRepository.UpdateUserEmail")
	dbConn := db.GetDatabaseInstance()

	newUser := models.Users{
		Email: payload.Email,
	}

	result := dbConn.Model(&models.Users{}).Where("id = ?", payload.Id).Updates(&newUser)
	if result.Error != nil {
		logrus.Error("Error updating the details ")
		return nil, result.Error
	}

	return nil, nil
}

func (repo UserRepository) GetUserDetails(userId int) (*dto.GenericResponseDto, error) {
	logrus.Info("UserRepository.GetUserDetails")
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

func (repo UserRepository) DeleteUserDetails(userId int) (*dto.GenericResponseDto, error) {
	logrus.Info("UserRepository.DeleteUserDetails")
	dbConn := db.GetDatabaseInstance()

	// Directly delete the record
	res := dbConn.Delete(&models.Users{}, "id = ?", userId)
	if res.Error != nil {
		logrus.WithError(res.Error).Error("Error deleting the record")
		return nil, errors.New("error deleting the record")
	}

	if res.RowsAffected == 0 {
		logrus.Warn("No record found to delete")
		return nil, errors.New("no record found with the given ID")
	}
	return nil, nil
}
