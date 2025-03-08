package constants

const (
	ErrEmailAlreadyExists     = "email already exists"
	ErrInvalidDateOfBirth     = "date of birth must be in YYYY-MM-DD format"
	ErrAgeLimitExceeded       = "date of birth must not be older than 100 years"
	ErrUserNotFound           = "user not found"
	ErrFailedToHashPassword   = "failed to hash password"
	ErrFailedToGenerateToken  = "failed to generate token"
	ErrInternalServerError    = "something went wrong"
	ErrInvalidCredentials     = "invalid email or password"
	ErrInvalidToken           = "invalid token"
	ErrTokenExpired           = "token expired"
	ErrInvalidPassword        = "invalid password"
	ErrInvalidOldPassword     = "incorrect old password"
	ErrRecordNotFound         = "record not found"
	ErrUnauthorized           = "not authorized or token expired"
	ErrFailedToCommit         = "failed to commit transaction"
	ErrFailedToFollow         = "failed to follow user"
	ErrFailedToCheckFollowing = "failed to check user following"
	ErrFailedToUpdate         = "failed to update user"
	ErrFailedToRetrieveUser   = "failed to retrieve user"
	ErrFailedToDeleteUser     = "failed to delete user"
	ErrFailedToDeleteResAddr  = "failed to delete residential details"
	ErrFailedToDeleteOffAddr  = "failed to delete office details"
	ErrFailedToUnfollow       = "failed to unfollow user"
	ErrFailedToUpdatePass     = "failed to update password"
	ErrUserCantFollowItself   = "user cannot follow themselves"
	ErrOldPassCantBeNewPass   = "new password can't be same as new password"

	ErrInvalidMaritalStatus = "invalid marital status value"
	ErrInvalidGenderValue   = "invalid gender value"

	SuccessLogOut          = "Successfully logged out"
	SuccessUserDeleted     = "User deleted successfully"
	SuccessUserFollowed    = "Users followed successfully"
	SuccessUserUnfollowed  = "Users unfollowed successfully"
	SuccessPasswordUpdated = "Password updates successfully"

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

	GenderMale   = 1
	GenderFemale = 2
	GenderOther  = 3

	MaritalStatusSingle  = 1
	MaritalStatusMarried = 2
)
