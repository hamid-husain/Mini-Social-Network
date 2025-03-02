package controllers

import (
	"mini-social-network/models"
	"mini-social-network/serializers"
	"mini-social-network/services"
	"mini-social-network/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var req serializers.SignUpRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:         req.Email,
		FirstName:     req.UserDetails.FirstName,
		LastName:      req.UserDetails.LastName,
		DateOfBirth:   req.UserDetails.DateOfBirth,
		Gender:        req.UserDetails.Gender,
		MaritalStatus: req.UserDetails.MaritalStatus,
		Password:      req.Password,
	}

	if err := services.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	office := models.OfficeDetail{
		UserID:        user.ID,
		EmployeeCode:  req.UserDetails.OfficeDetails.EmployeeCode,
		Address:       req.UserDetails.OfficeDetails.Address,
		City:          req.UserDetails.OfficeDetails.City,
		State:         req.UserDetails.OfficeDetails.State,
		Country:       req.UserDetails.OfficeDetails.Country,
		ContactNumber: req.UserDetails.OfficeDetails.ContactNo,
		Email:         req.UserDetails.OfficeDetails.Email,
		Name:          req.UserDetails.OfficeDetails.Name,
	}

	if err := services.SaveOfficeDetails(&office); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resident := models.ResidentialDetail{
		UserID:         user.ID,
		Address:        req.UserDetails.ResidentialDetails.Address,
		City:           req.UserDetails.ResidentialDetails.City,
		State:          req.UserDetails.ResidentialDetails.State,
		Country:        req.UserDetails.ResidentialDetails.Country,
		ContactNumber1: req.UserDetails.ResidentialDetails.ContactNo1,
		ContactNumber2: req.UserDetails.ResidentialDetails.ContactNo2,
	}

	if err := services.SaveResidentialDetail(&resident); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var officeDetails []models.OfficeDetail
	var residentialDetails []models.ResidentialDetail

	officeDetails = append(officeDetails, office)
	residentialDetails = append(residentialDetails, resident)

	if err := services.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user with office and residential details"})
		return
	}

	token, expiryTime, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	userResponse := serializers.SerializeResponse(user, resident, office)
	tokenResponse := serializers.SerializeToken(token, expiryTime)
	c.JSON(http.StatusCreated, gin.H{
		"id":            user.ID,
		"user_id":       user.ID,
		"email":         user.Email,
		"last_modified": user.UpdatedAt,
		"user_details":  userResponse,
		"token":         tokenResponse,
	})
}
