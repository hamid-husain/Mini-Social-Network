package utils

import (
	"errors"
	"mini-social-network/config"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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
	errorMessages := make(map[string]string)

	errorLines := strings.Split(err.Error(), "\n")

	for _, line := range errorLines {
		switch {
		case strings.Contains(line, "Email") && strings.Contains(line, "'email' tag"):
			errorMessages["email"] = "Please enter a valid email address"
		case strings.Contains(line, "Password") && strings.Contains(line, "'min' tag"):
			errorMessages["password"] = "Password must be at least 6 characters long"
		case strings.Contains(line, "UserDetails.FirstName"):
			errorMessages["first_name"] = "First name is required"
		case strings.Contains(line, "UserDetails.DateOfBirth"):
			errorMessages["date_of_birth"] = "Date of birth is required"
		case strings.Contains(line, "UserDetails.Gender") && strings.Contains(line, "'oneof' tag"):
			errorMessages["gender"] = "Gender must be 1 (male), 2 (female), or 3 (other)"
		case strings.Contains(line, "UserDetails.MaritalStatus") && strings.Contains(line, "'oneof' tag"):
			errorMessages["marital_status"] = "Marital status must be 1 (single) or 2 (married)"
		case strings.Contains(line, "UserDetails.ResidentialDetails.Address"):
			errorMessages["residential_address"] = "Residential address is required"
		case strings.Contains(line, "UserDetails.ResidentialDetails.City"):
			errorMessages["residential_city"] = "Residential city is required"
		case strings.Contains(line, "UserDetails.ResidentialDetails.State"):
			errorMessages["residential_state"] = "Residential state is required"
		case strings.Contains(line, "UserDetails.ResidentialDetails.Country"):
			errorMessages["residential_country"] = "Residential country is required"
		case strings.Contains(line, "UserDetails.ResidentialDetails.ContactNo1") && strings.Contains(line, "'e164' tag"):
			errorMessages["residential_contact_1"] = "Primary contact number must be in valid E.164 format"
		case strings.Contains(line, "UserDetails.ResidentialDetails.ContactNo2") && strings.Contains(line, "'e164' tag"):
			errorMessages["residential_contact_2"] = "Secondary contact number must be in valid E.164 format"
		case strings.Contains(line, "UserDetails.OfficeDetails.EmployeeCode"):
			errorMessages["employee_code"] = "Employee code is required"
		case strings.Contains(line, "UserDetails.OfficeDetails.Address"):
			errorMessages["office_address"] = "Office address is required"
		case strings.Contains(line, "UserDetails.OfficeDetails.City"):
			errorMessages["office_city"] = "Office city is required"
		case strings.Contains(line, "UserDetails.OfficeDetails.State"):
			errorMessages["office_state"] = "Office state is required"
		case strings.Contains(line, "UserDetails.OfficeDetails.Country"):
			errorMessages["office_country"] = "Office country is required"
		case strings.Contains(line, "UserDetails.OfficeDetails.ContactNo"):
			errorMessages["office_contact"] = "Office contact number is required"
		case strings.Contains(line, "UserDetails.OfficeDetails.Email") && strings.Contains(line, "'email' tag"):
			errorMessages["office_email"] = "Please enter a valid office email address"
		case strings.Contains(line, "UserDetails.OfficeDetails.Name"):
			errorMessages["office_name"] = "Office name is required"
		default:
			errorMessages["error"] = "Invalid input data"
		}
	}

	return gin.H{"error:": errorMessages}
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
