package serializers

import (
	"mini-social-network/models/users"
)

type UserDetailsInput struct {
	FirstName          string                  `json:"first_name" binding:"required"`
	LastName           string                  `json:"last_name"`
	DateOfBirth        string                  `json:"date_of_birth" binding:"required"`
	Gender             uint8                   `json:"gender" binding:"omitempty,oneof=0 1 2"`
	MaritalStatus      uint8                   `json:"marital_status" binding:"oneof=0 1"`
	ResidentialDetails ResidentialDetailsInput `json:"residential_details"`
	OfficeDetails      OfficeDetailsInput      `json:"office_details"`
}

type userDetailsResponse struct {
	ID                 uint                       `json:"id"`
	Email              string                     `json:"email"`
	FirstName          string                     `json:"first_name"`
	LastName           string                     `json:"last_name,omitempty"`
	DateOfBirth        string                     `json:"date_of_birth"`
	Gender             uint8                      `json:"gender"`
	MaritalStatus      uint8                      `json:"marital_status"`
	ResidentialDetails ResidentialDetailsResponse `json:"residential_details"`
	OfficeDetails      OfficeDetailsResponse      `json:"office_details"`
}

func SerializeResponse(user users.User, resident users.ResidentialDetail, office users.OfficeDetail) userDetailsResponse {
	return userDetailsResponse{
		ID:            user.ID,
		Email:         user.Email,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		DateOfBirth:   user.DateOfBirth,
		Gender:        user.Gender,
		MaritalStatus: user.MaritalStatus,
		ResidentialDetails: ResidentialDetailsResponse{
			ID:         resident.ID,
			UserID:     resident.UserID,
			Address:    resident.Address,
			City:       resident.City,
			State:      resident.State,
			Country:    resident.Country,
			ContactNo1: resident.ContactNumber1,
			ContactNo2: resident.ContactNumber2,
		},
		OfficeDetails: OfficeDetailsResponse{
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
		},
	}
}
