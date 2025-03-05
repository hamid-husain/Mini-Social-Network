package serializers

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	Token TokenSerializer   `json:"token"`
	User  UserLoginResponse `json:"user"`
}

func SerializeLoginResponse(user UserLoginResponse, token TokenSerializer) LoginResponse {
	return LoginResponse{
		User:  user,
		Token: token,
	}
}
