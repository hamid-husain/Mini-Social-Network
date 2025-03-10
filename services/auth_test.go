package services_test

import (
	"github.com/stretchr/testify/assert"

	"errors"
	"testing"

	"mini-social-network/constants"
	"mini-social-network/mocks"
	"mini-social-network/models"
	"mini-social-network/serializers"
)

func TestCreateUserWithDetails(t *testing.T) {
	signUpReq := &serializers.SignUpRequest{
		Email:    "test@example.com",
		Password: "password123",
		UserDetails: serializers.UserDetailsInput{
			FirstName:     "John",
			LastName:      "Doe",
			DateOfBirth:   "1990-01-01",
			Gender:        "Male",
			MaritalStatus: "Single",
			ResidentialDetails: serializers.ResidentialDetailsInput{
				Address:    "123 Main St",
				City:       "City",
				State:      "State",
				Country:    "Country",
				ContactNo1: "123456789",
				ContactNo2: "987654321",
			},
			OfficeDetails: serializers.OfficeDetailsInput{
				EmployeeCode: "EMP001",
				Address:      "Office Address",
				City:         "City",
				State:        "State",
				Country:      "Country",
				ContactNo:    "123456789",
				Email:        "office@example.com",
				Name:         "Company",
			},
		},
	}

	mockAuthService := mocks.NewAuthServiceInt(t)
	mockAuthService.On("CreateUserWithDetails", signUpReq).Return(&serializers.SignUpResponse{
		UserDetails:   serializers.UserDetailsResponse{Email: "test@example.com"},
		TokenResponse: serializers.TokenSerializer{Key: "mockToken", ExpiryTime: 3600},
	}, nil)

	response, err := mockAuthService.CreateUserWithDetails(signUpReq)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "test@example.com", response.UserDetails.Email)
	mockAuthService.AssertExpectations(t)
}

func TestLoginHandler_Success(t *testing.T) {
	loginReq := serializers.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	mockUser := &models.User{
		Email:    "test@example.com",
		Password: "hashedpassword",
	}

	mockAuthService := mocks.NewAuthServiceInt(t)
	mockAuthService.On("VerifyUserCredentials", loginReq.Email, loginReq.Password).Return(mockUser, nil)

	mockToken := "mockToken"
	mockExpiryTime := 3600
	mockAuthService.On("GenerateJWT", mockUser.ID).Return(mockToken, mockExpiryTime, nil)

	mockAuthService.On("LoginHandler", loginReq).Return(&serializers.LoginResponse{
		User:  serializers.SerializeUserLoginResponse(*mockUser),
		Token: serializers.TokenSerializer{Key: mockToken, ExpiryTime: int64(mockExpiryTime)},
	}, nil)

	response, err := mockAuthService.LoginHandler(loginReq)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "test@example.com", response.User.Email)
	assert.Equal(t, mockToken, response.Token.Key)
	mockAuthService.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	user := &models.User{
		Email:    "test@example.com",
		Password: "hashedpassword",
	}

	mockAuthService := mocks.NewAuthServiceInt(t)
	mockAuthService.On("UpdateUser", user).Return(nil)

	err := mockAuthService.UpdateUser(user)

	assert.NoError(t, err)
	mockAuthService.AssertExpectations(t)
}

func TestVerifyUserCredentials_Failure(t *testing.T) {
	loginReq := serializers.LoginRequest{
		Email:    "test@example.com",
		Password: "wrongpassword",
	}

	mockAuthService := mocks.NewAuthServiceInt(t)
	mockAuthService.On("VerifyUserCredentials", loginReq.Email, loginReq.Password).Return(nil, errors.New(constants.ErrInvalidPassword))

	_, err := mockAuthService.LoginHandler(loginReq)

	assert.Error(t, err)
	mockAuthService.AssertExpectations(t)
}

func TestVerifyUserCredentials_Success(t *testing.T) {
	email := "test@example.com"
	password := "password123"

	mockUser := &models.User{
		Email:    email,
		Password: "hashedpassword",
	}

	mockAuthService := mocks.NewAuthServiceInt(t)
	mockAuthService.On("VerifyUserCredentials", email, password).Return(mockUser, nil)

	user, err := mockAuthService.VerifyUserCredentials(email, password)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, email, user.Email)
	mockAuthService.AssertExpectations(t)
}
