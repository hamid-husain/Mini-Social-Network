package services

import (
	"errors"
	"mini-social-network/db"
	"mini-social-network/models"
	"mini-social-network/utils"

	"gorm.io/gorm"
)

func CreateUser(user *models.User) error {
	var existingUser models.User
	if err := db.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New("email already in use")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = hashedPassword

	if err := db.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func FindUsers() ([]models.User, error) {
	var userList []models.User
	result := db.DB.Find(&userList)
	return userList, result.Error
}

func UpdateUser(user *models.User) error {
	if err := db.DB.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func SaveOfficeDetails(office *models.OfficeDetail) error {
	if err := db.DB.Create(office).Error; err != nil {
		return err
	}

	return nil
}

func SaveResidentialDetail(resident *models.ResidentialDetail) error {
	if err := db.DB.Create(resident).Error; err != nil {
		return err
	}

	return nil
}
