package services

import (
	"github.com/rabinsharma25/ims-backend/src/dto"
	"github.com/rabinsharma25/ims-backend/src/repositories"
	"github.com/sirupsen/logrus"
)

type UserService struct{}

var userRepository repositories.UserRepository = repositories.UserRepository{}

func (service UserService) AddUserDetails(payload dto.AddUserDetailsReqDto) (*dto.GenericResponseDto, error) {
	logrus.Info("UserService.AddUserDetails")

	_, err := userRepository.AddUserDetails(payload)

	if err != nil {
		logrus.Error("Failed to Add user details ")
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
		Message: "Successfully added the customer details",
		Data:    nil,
	}, nil

}

func (service UserService) UpdateUserDetails(payload dto.UpdateUserDetailsReqDto) (*dto.GenericResponseDto, error) {
	logrus.Info("UserService.UpdateUserDetails")

	_, err := userRepository.UpdateUserDetails(payload)

	if err != nil {
		logrus.Error("Failed to update customer details ")
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
		Message: "Successfully updated user details",
		Data:    nil,
	}, nil

}

func (service UserService) UpdateUserEmail(payload dto.UpdateUserEmailReqDto) (*dto.GenericResponseDto, error) {
	logrus.Info("UserService.UpdateUserEmail")

	_, err := userRepository.UpdateUserEmail(payload)

	if err != nil {
		logrus.Error("Failed to update customer email ")
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
		Message: "Successfully updated user email",
		Data:    nil,
	}, nil

}

func (service UserService) GetUserDetails(userId int) (*dto.GenericResponseDto, error) {
	logrus.Info("UserService.GetUserDetails")

	res, err := userRepository.GetUserDetails(userId)

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
		Message: "Successfully fetched store owner details",
		Data:    res.Data,
	}, nil

}
func (service UserService) DeleteUserDetails(userId int) (*dto.GenericResponseDto, error) {
	logrus.Info("UserService.DeleteUserDetails")

	_, err := userRepository.DeleteUserDetails(userId)

	if err != nil {
		logrus.Error("Failed to delete user details ")
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
		Message: "Successfully deleted user details",
		Data:    nil,
	}, nil

}
