package services

import (
	"mini-social-network/db"
	"mini-social-network/models"
	"mini-social-network/serializers"
)

func DeleteUserByID(userID string) (*serializers.DeleteUserResponse, error) {
	tx := db.DB.Begin()

	var user models.User
	if err := tx.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	if err := tx.Delete(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var officeDetails []models.OfficeDetail
	if err := tx.Where("user_id = ?", user.ID).Unscoped().Find(&officeDetails).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Unscoped().Delete(&officeDetails).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var residentialDetails []models.ResidentialDetail
	if err := tx.Where("user_id = ?", user.ID).Unscoped().Find(&residentialDetails).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Unscoped().Delete(&residentialDetails).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	userResponse := serializers.SerializeResponse(user, residentialDetails, officeDetails)
	response := serializers.SerializeDeleteUserResponse(user, userResponse)

	return &response, nil
}
