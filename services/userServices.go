package services

import (
	"errors"
	"mini-social-network/db"
	"mini-social-network/models/users"
	"mini-social-network/utils"

	"gorm.io/gorm"
)

func CreateUser(user *users.User) error {
	var existingUser users.User
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

func GetUserByEmail(email string) (*users.User, error) {
	var user users.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func FindUsers() ([]users.User, error) {
	var userList []users.User
	result := db.DB.Find(&userList)
	return userList, result.Error
}

func UpdateUser(user *users.User) error {
	if err := db.DB.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func SaveOfficeDetails(office *users.OfficeDetail) error {
	if err := db.DB.Create(office).Error; err != nil {
		return err
	}

	return nil
}

func SaveResidentialDetail(resident *users.ResidentialDetail) error {
	if err := db.DB.Create(resident).Error; err != nil {
		return err
	}

	return nil
}
