package utils

import (
	"errors"
	"mini-social-network/config"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

var jwtSecret = []byte(config.AppConfig.SecretKey)

func GenerateJWT(userID uint) (string, int64, error) {
	expiryTime := time.Now().Add(4 * time.Hour).Unix()

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expiryTime,
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", 0, err
	}

	return signedToken, expiryTime, nil
}

func ParseValidationErrors(err error) gin.H {
	validationErrors := make(map[string][]string)

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			fieldName := e.Field()

			var errorMessage string
			switch e.Tag() {
			case "required":
				errorMessage = fieldName + " cannot be left blank."
			case "email":
				errorMessage = "Please provide a valid email address."
			case "min":
				errorMessage = fieldName + " must have at least " + e.Param() + " characters."
			case "max":
				errorMessage = fieldName + " cannot exceed " + e.Param() + " characters."
			case "oneof":
				errorMessage = fieldName + " should be one of the following options: " + e.Param() + "."
			case "e164":
				errorMessage = fieldName + " should follow the international phone format."
			default:
				errorMessage = fieldName + " is invalid."
			}

			validationErrors[fieldName] = append(validationErrors[fieldName], errorMessage)
		}
	}

	return gin.H{"error:": validationErrors}
}

const MaxAgeLimit = 100

func ValidateDOB(dobStr string) error {
	const layout = "2006-01-02"

	dob, err := time.Parse(layout, dobStr)
	if err != nil {
		return errors.New("date of birth must be in YYYY-MM-DD format")
	}

	today := time.Now()
	age := today.Year() - dob.Year()

	if today.YearDay() < dob.YearDay() {
		age--
	}

	if age > MaxAgeLimit {
		return errors.New("date of birth must not be older than 100 years")
	}

	return nil
}
