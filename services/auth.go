package services

import (
	"gorm.io/gorm"

	"errors"

	"mini-social-network/constants"
	"mini-social-network/db"
	"mini-social-network/models"
	"mini-social-network/serializers"
	"mini-social-network/utils"
)

func CreateUserWithDetails(req *serializers.SignUpRequest) (*serializers.SignUpResponse, error) {
	tx := db.DB.Begin()

	user := models.User{
		Email:         req.Email,
		FirstName:     req.UserDetails.FirstName,
		LastName:      req.UserDetails.LastName,
		DateOfBirth:   req.UserDetails.DateOfBirth,
		Gender:        req.UserDetails.Gender,
		MaritalStatus: req.UserDetails.MaritalStatus,
		Password:      req.Password,
	}

	if err := CreateUser(tx, &user); err != nil {
		tx.Rollback()
		return nil, err
	}

	office := models.OfficeDetail{
		UserID:        user.ID,
		EmployeeCode:  req.UserDetails.OfficeDetails.EmployeeCode,
		Address:       req.UserDetails.OfficeDetails.Address,
		City:          req.UserDetails.OfficeDetails.City,
		State:         req.UserDetails.OfficeDetails.State,
		Country:       req.UserDetails.OfficeDetails.Country,
		ContactNumber: req.UserDetails.OfficeDetails.ContactNo,
		Email:         req.UserDetails.OfficeDetails.Email,
		Name:          req.UserDetails.OfficeDetails.Name,
	}

	if err := SaveOfficeDetails(tx, &office); err != nil {
		tx.Rollback()
		return nil, err
	}

	resident := models.ResidentialDetail{
		UserID:         user.ID,
		Address:        req.UserDetails.ResidentialDetails.Address,
		City:           req.UserDetails.ResidentialDetails.City,
		State:          req.UserDetails.ResidentialDetails.State,
		Country:        req.UserDetails.ResidentialDetails.Country,
		ContactNumber1: req.UserDetails.ResidentialDetails.ContactNo1,
		ContactNumber2: req.UserDetails.ResidentialDetails.ContactNo2,
	}

	if err := SaveResidentialDetail(tx, &resident); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	token, expiryTime, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	userResponse := serializers.SerializeResponse(user, resident, office)
	tokenResponse := serializers.SerializeToken(token, expiryTime)

	response := serializers.SerializeSignUpResponse(user, userResponse, tokenResponse)

	return &response, nil
}

func CreateUser(tx *gorm.DB, user *models.User) error {
	var existingUser models.User
	if err := tx.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New(constants.ErrEmailAlreadyExists)
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return errors.New(constants.ErrFailedToHashPassword)
	}
	user.Password = hashedPassword

	if err := tx.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.ErrUserNotFound)
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

func SaveOfficeDetails(tx *gorm.DB, office *models.OfficeDetail) error {
	if err := tx.Create(office).Error; err != nil {
		return err
	}

	return nil
}

func SaveResidentialDetail(tx *gorm.DB, resident *models.ResidentialDetail) error {
	if err := tx.Create(resident).Error; err != nil {
		return err
	}

	return nil
}
