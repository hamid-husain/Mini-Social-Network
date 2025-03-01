package utils

import (
	"mini-social-network/config"
	"time"

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
