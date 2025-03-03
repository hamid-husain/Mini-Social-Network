package serializers

import (
	"mini-social-network/models"
	"time"
)

type SignUpRequest struct {
	Email       string           `json:"email" binding:"required,email"`
	Password    string           `json:"password" binding:"required,min=6"`
	UserDetails UserDetailsInput `json:"user_details" binding:"required"`
}

type SignUpResponse struct {
	ID            uint                `json:"id"`
	UserID        uint                `json:"user_id"`
	Email         string              `json:"email"`
	UpdatedAt     time.Time           `json:"last_modified"`
	UserDetails   userDetailsResponse `json:"user_details"`
	TokenResponse TokenSerializer     `json:"token"`
}

func SerializeSignUpResponse(user models.User, userResponse userDetailsResponse, token TokenSerializer) SignUpResponse {
	return SignUpResponse{
		ID:            user.ID,
		UserID:        user.ID,
		Email:         user.Email,
		UpdatedAt:     user.UpdatedAt,
		UserDetails:   userResponse,
		TokenResponse: token,
	}
}
