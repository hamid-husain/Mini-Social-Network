package serializers

import "mini-social-network/models"

type UserFollowRequest struct {
	UserIDs []uint `json:"user_ids"`
}

type FollowResponse struct {
	ID            uint   `json:"id"`
	Email         string `json:"email"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name,omitempty"`
	DateOfBirth   string `json:"date_of_birth"`
	Gender        string `json:"gender"`
	MaritalStatus string `json:"marital_status"`
}

func SerializeFollowResponse(users []models.User) []FollowResponse {
	var userResponses []FollowResponse
	for _, user := range users {
		userResponses = append(userResponses, FollowResponse{
			ID:            user.ID,
			Email:         user.Email,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			DateOfBirth:   user.DateOfBirth,
			Gender:        user.GetGender(),
			MaritalStatus: user.GetGender(),
		})
	}
	return userResponses
}
