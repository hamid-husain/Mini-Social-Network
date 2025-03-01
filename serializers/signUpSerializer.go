package serializers

type SignUpRequest struct {
	Email       string           `json:"email" binding:"required,email"`
	Password    string           `json:"password" binding:"required,min=6"`
	UserDetails UserDetailsInput `json:"user_details" binding:"required"`
}
