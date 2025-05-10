package constants

const (
	ErrEmailAlreadyExists    = "email already exists"
	ErrInvalidDateOfBirth    = "date of birth must be in YYYY-MM-DD format"
	ErrAgeLimitExceeded      = "date of birth must not be older than 100 years"
	ErrUserNotFound          = "user not found"
	ErrFailedToHashPassword  = "failed to hash password"
	ErrFailedToGenerateToken = "failed to generate token"
	ErrInternalServerError   = "internal server error"
)
