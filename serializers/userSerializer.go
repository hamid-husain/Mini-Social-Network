package serializers

import (
	"mini-social-network/models/users"
)

type CreateUser struct {
	Email         string `json:"email" binding:"required,email"`
	FirstName     string `json:"first_name" binding:"required"`
	LastName      string `json:"last_name"`
	DateOfBirth   string `json:"date_of_birth" binding:"required"`
	Gender        uint8  `json:"gender" binding:"omitempty,oneof=0 1 2"`
	MaritalStatus uint8  `json:"marital_status" binding:"omitempty,oneof=0 1"`
	Password      string `json:"password" binding:"required,min=6"`
}

type UserResponse struct {
	ID            uint   `json:"id"`
	Email         string `json:"email"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name,omitempty"`
	DateOfBirth   string `json:"date_of_birth"`
	Gender        uint8  `json:"gender"`
	MaritalStatus uint8  `json:"marital_status"`
}

func SerializeUser(user users.User) UserResponse {
	return UserResponse{
		ID:            user.ID,
		Email:         user.Email,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		DateOfBirth:   user.DateOfBirth,
		Gender:        user.Gender,
		MaritalStatus: user.MaritalStatus,
	}
}
