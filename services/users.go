package services

import (
	"errors"
	"mini-social-network/constants"
	"mini-social-network/models"
	"mini-social-network/serializers"
	"time"

	"gorm.io/gorm"
)

func (s *Service) GetUserByID(userID uint) (*serializers.GetUserResponse, error) {
	var user models.User
	if err := s.DB.Where("id = ? AND deleted_at IS NULL", userID).First(&user).Error; err != nil {
		return nil, err
	}

	var officeDetails []models.OfficeDetail
	if err := s.DB.Where("user_id = ?", userID).Find(&officeDetails).Error; err != nil {
		return nil, err
	}

	var residentialDetails []models.ResidentialDetail
	if err := s.DB.Where("user_id = ?", userID).Find(&residentialDetails).Error; err != nil {
		return nil, err
	}

	userResponse := serializers.SerializeResponse(user, residentialDetails, officeDetails)
	response := serializers.SerializeGetUserResponse(user, userResponse)

	return &response, nil
}

func (s *Service) ListUsers() ([]serializers.ListUserResponse, error) {
	var users []models.User

	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	response := serializers.SerializeListUser(users)

	return response, nil
}

func (s *Service) DeleteUserByID(userID uint) (*serializers.DeleteUserResponse, error) {
	tx := s.DB.Begin()
	var user models.User

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.First(&user, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return nil, errors.New(constants.ErrUserNotFound)
		}
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&user).Update("deleted_at", time.Now()).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var officeDetails []models.OfficeDetail
	if err := tx.Where("user_id = ?", userID).Find(&officeDetails).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, office := range officeDetails {
		if err := tx.Model(&office).Update("deleted_at", time.Now()).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	var residentialDetails []models.ResidentialDetail
	if err := tx.Where("user_id = ?", userID).Find(&residentialDetails).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, resident := range residentialDetails {
		if err := tx.Model(&resident).Update("deleted_at", time.Now()).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	userResponse := serializers.SerializeResponse(user, residentialDetails, officeDetails)
	response := serializers.SerializeDeleteUserResponse(user, userResponse)

	return &response, nil
}

func (s *Service) UpdateUserByID(userID uint, req *serializers.UpdateUserRequest) (*serializers.GetUserResponse, error) {
	tx := s.DB.Begin()

	var user models.User
	if err := tx.Where("id = ?", userID).First(&user).Error; err != nil {
		tx.Rollback()
		return nil, errors.New(constants.ErrUserNotFound)
	}

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Gender != "" {
		var gender uint8
		var err error

		switch req.Gender {
		case "male":
			gender = 1
		case "female":
			gender = 2
		case "other":
			gender = 3
		default:
			err = errors.New("invalid gender value")
		}

		if err != nil {
			tx.Rollback()
			return nil, err
		}
		user.Gender = gender
	}

	if req.DateOfBirth != "" {
		user.DateOfBirth = req.DateOfBirth
	}
	if req.MaritalStatus != "" {
		var maritalStatus uint8
		var err error

		switch req.MaritalStatus {
		case "single":
			maritalStatus = 1
		case "married":
			maritalStatus = 2
		default:
			err = errors.New("invalid marital status value")
		}

		if err != nil {
			tx.Rollback()
			return nil, err
		}
		user.MaritalStatus = maritalStatus
	}

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	var officeDetails []models.OfficeDetail
	if err := s.DB.Where("user_id = ?", userID).Find(&officeDetails).Error; err != nil {
		return nil, err
	}

	var residentialDetails []models.ResidentialDetail
	if err := s.DB.Where("user_id = ?", userID).Find(&residentialDetails).Error; err != nil {
		return nil, err
	}

	userResponse := serializers.SerializeResponse(user, residentialDetails, officeDetails)
	response := serializers.SerializeGetUserResponse(user, userResponse)

	return &response, nil
}
