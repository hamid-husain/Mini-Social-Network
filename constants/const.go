package constants

const (
	ErrEmailAlreadyExists    = "email already exists"
	ErrInvalidDateOfBirth    = "date of birth must be in YYYY-MM-DD format"
	ErrAgeLimitExceeded      = "date of birth must not be older than 100 years"
	ErrUserNotFound          = "user not found"
	ErrFailedToHashPassword  = "failed to hash password"
	ErrFailedToGenerateToken = "failed to generate token"
	ErrInternalServerError   = "internal server error"
	ErrInvalidCredentials    = "invalid email or password"
	ErrInvalidToken          = "invalid token"
	ErrTokenExpired          = "token expired"
	ErrInvalidPassword       = "invalid password"
	ErrRecordNotFound        = "record not found"
	ErrUnauthorized          = "access unauthorized"

	ErrInvalidMaritalStatus = "invalid marital status value"
	ErrInvalidGenderValue   = "invalid gender value"

	SuccessLogOut = "Successfully logged out"

	DBHost       = "DB_HOST"
	DBPort       = "DB_PORT"
	DBUser       = "DB_USER"
	DBPassword   = "DB_PASSWORD"
	DBName       = "DB_NAME"
	DBSSLMode    = "DB_SSL"
	ServerPort   = "SERVER_PORT"
	SecretKey    = "SECRET_KEY"
	Domain       = "DOMAIN"
	MigrationDir = "MIGRATION_DIR"

	Male    = "male"
	Female  = "female"
	Other   = "other"
	Unknown = "unknown"
	Single  = "single"
	Married = "married"

	Authorization = "Authorization"
)
