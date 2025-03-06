package services

import (
	"net/http"

	"gorm.io/gorm"

	"errors"

	"mini-social-network/constants"
	apiError "mini-social-network/errors"
	"mini-social-network/models"
	"mini-social-network/serializers"
	"mini-social-network/utils"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) CreateUserWithDetails(req *serializers.SignUpRequest) (*serializers.SignUpResponse, error) {
	tx := s.DB.Begin()

	var gender, maritalStatus uint8
	var err error

	switch req.UserDetails.Gender {
	case constants.Male:
		gender = 1
	case constants.Female:
		gender = 2
	case constants.Other:
		gender = 3
	default:
		err = errors.New(constants.ErrInvalidGenderValue)
	}

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	switch req.UserDetails.MaritalStatus {
	case constants.Single:
		maritalStatus = 1
	case constants.Married:
		maritalStatus = 2
	default:
		err = errors.New(constants.ErrInvalidMaritalStatus)
	}

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	user := models.User{
		Email:         req.Email,
		FirstName:     req.UserDetails.FirstName,
		LastName:      req.UserDetails.LastName,
		DateOfBirth:   req.UserDetails.DateOfBirth,
		Gender:        gender,
		MaritalStatus: maritalStatus,
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

	bearerToken := "Bearer " + token

	userResponse := serializers.SerializeResponse(user, []models.ResidentialDetail{resident}, []models.OfficeDetail{office})
	tokenResponse := serializers.SerializeToken(bearerToken, expiryTime)

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

func (s *Service) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.ErrUserNotFound)
		}
		return nil, err
	}
	return &user, nil
}

func (s *Service) FindUsers() ([]models.User, error) {
	var userList []models.User
	result := s.DB.Find(&userList)
	return userList, result.Error
}

func (s *Service) UpdateUser(user *models.User) error {
	if err := s.DB.Save(user).Error; err != nil {
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

func (s *Service) LoginHandler(req serializers.LoginRequest) (*serializers.LoginResponse, *apiError.APIError) {
	user, err := s.VerifyUserCredentials(req.Email, req.Password)
	if err != nil {
		return nil, apiError.NewAPIError(constants.ErrInvalidCredentials, http.StatusUnauthorized)
	}

	token, expiryTime, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, apiError.NewAPIError(constants.ErrFailedToGenerateToken, http.StatusInternalServerError)
	}

	bearerToken := "Bearer " + token

	userResponse := serializers.SerializeUserLoginResponse(*user)
	tokenResponse := serializers.SerializeToken(bearerToken, expiryTime)
	response := serializers.SerializeLoginResponse(userResponse, tokenResponse)

	return &response, nil
}

func (s *Service) VerifyUserCredentials(email, password string) (*models.User, error) {
	var user models.User
	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if err.Error() == constants.ErrRecordNotFound {
			return nil, errors.New(constants.ErrUserNotFound)
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New(constants.ErrInvalidPassword)
	}

	return &user, nil
}
