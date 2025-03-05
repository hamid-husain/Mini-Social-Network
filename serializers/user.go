package serializers

import (
	"mini-social-network/constants"
	"mini-social-network/models"
	"time"
)

type UserDetailsInput struct {
	FirstName          string                  `json:"first_name" binding:"required"`
	LastName           string                  `json:"last_name"`
	DateOfBirth        string                  `json:"date_of_birth" binding:"required,valid_dob"`
	Gender             string                  `json:"gender" binding:"required, oneof=male female other"`
	MaritalStatus      string                  `json:"marital_status" binding:"required, oneof=single married"`
	ResidentialDetails ResidentialDetailsInput `json:"residential_details"`
	OfficeDetails      OfficeDetailsInput      `json:"office_details"`
}

type UserDetailsResponse struct {
	ID                 uint                         `json:"id"`
	Email              string                       `json:"email"`
	FirstName          string                       `json:"first_name"`
	LastName           string                       `json:"last_name,omitempty"`
	DateOfBirth        string                       `json:"date_of_birth"`
	Gender             string                       `json:"gender"`
	MaritalStatus      string                       `json:"marital_status"`
	ResidentialDetails []ResidentialDetailsResponse `json:"residential_details"`
	OfficeDetails      []OfficeDetailsResponse      `json:"office_details"`
}

type UserLoginResponse struct {
	ID           uint   `json:"id"`
	Email        string `json:"email"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	LastModified string `json:"last_modified"`
}

type DeleteUserResponse struct {
	Message string              `json:"message"`
	User    UserDetailsResponse `json:"user"`
}

type GetUserResponse struct {
	ID          uint                `json:"id"`
	UserID      uint                `json:"user_id"`
	Email       string              `json:"email"`
	UpdatedAt   time.Time           `json:"last_modified"`
	UserDetails UserDetailsResponse `json:"user_details"`
}

func SerializeUserLoginResponse(user models.User) UserLoginResponse {
	return UserLoginResponse{
		ID:           user.ID,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		LastModified: user.UpdatedAt.String(),
	}
}
func SerializeResponse(user models.User, residents []models.ResidentialDetail, offices []models.OfficeDetail) UserDetailsResponse {
	var residentialDetails []ResidentialDetailsResponse
	for _, resident := range residents {
		residentialDetails = append(residentialDetails, ResidentialDetailsResponse{
			ID:         resident.ID,
			UserID:     resident.UserID,
			Address:    resident.Address,
			City:       resident.City,
			State:      resident.State,
			Country:    resident.Country,
			ContactNo1: resident.ContactNumber1,
			ContactNo2: resident.ContactNumber2,
		})
	}

	var officeDetails []OfficeDetailsResponse
	for _, office := range offices {
		officeDetails = append(officeDetails, OfficeDetailsResponse{
			ID:           office.ID,
			UserID:       office.UserID,
			EmployeeCode: office.EmployeeCode,
			Address:      office.Address,
			City:         office.City,
			State:        office.State,
			Country:      office.Country,
			ContactNo:    office.ContactNumber,
			Email:        office.Email,
			Name:         office.Name,
		})
	}

	var gender string
	switch user.Gender {
	case 1:
		gender = constants.Male
	case 2:
		gender = constants.Female
	case 3:
		gender = constants.Other
	default:
		gender = constants.Unknown
	}

	var maritalStatus string
	switch user.MaritalStatus {
	case 1:
		maritalStatus = constants.Single
	case 2:
		maritalStatus = constants.Married
	default:
		maritalStatus = constants.Unknown
	}

	return UserDetailsResponse{
		ID:                 user.ID,
		Email:              user.Email,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		DateOfBirth:        user.DateOfBirth,
		Gender:             gender,
		MaritalStatus:      maritalStatus,
		ResidentialDetails: residentialDetails,
		OfficeDetails:      officeDetails,
	}
}

func SerializeDeleteUserResponse(user models.User, userResponse UserDetailsResponse) DeleteUserResponse {
	return DeleteUserResponse{
		Message: "User and associated details deleted successfully",
		User:    userResponse,
	}
}

func SerializeGetUserResponse(user models.User, userResponse UserDetailsResponse) GetUserResponse {
	return GetUserResponse{
		ID:          user.ID,
		UserID:      user.ID,
		Email:       user.Email,
		UpdatedAt:   user.UpdatedAt,
		UserDetails: userResponse,
	}
}
