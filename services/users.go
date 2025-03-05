package services

import (
	"errors"
	"mini-social-network/constants"
	"mini-social-network/db"
	"mini-social-network/models"
	"mini-social-network/serializers"
	"time"

	"gorm.io/gorm"
)

func GetUserByID(userID uint) (*serializers.GetUserResponse, error) {
	var user models.User
	if err := db.DB.Where("id = ? AND deleted_at IS NULL", userID).First(&user).Error; err != nil {
		return nil, err
	}

	var officeDetails []models.OfficeDetail
	if err := db.DB.Where("user_id = ?", userID).Find(&officeDetails).Error; err != nil {
		return nil, err
	}

	var residentialDetails []models.ResidentialDetail
	if err := db.DB.Where("user_id = ?", userID).Find(&residentialDetails).Error; err != nil {
		return nil, err
	}

	userResponse := serializers.SerializeResponse(user, residentialDetails, officeDetails)
	response := serializers.SerializeGetUserResponse(user, userResponse)

	return &response, nil
}

func DeleteUserByID(userID uint) (*serializers.DeleteUserResponse, error) {
	tx := db.DB.Begin()
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
