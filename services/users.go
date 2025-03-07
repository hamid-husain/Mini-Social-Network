package services

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"fmt"
	"net/http"

	"mini-social-network/constants"
	"mini-social-network/errors"
	"mini-social-network/models"
	"mini-social-network/serializers"
	"mini-social-network/utils"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) GetUserByID(userID uint) (*serializers.GetUserResponse, *errors.APIError) {
	var user models.User
	if err := s.DB.Preload("OfficeDetails").Preload("ResidentialDetails").
		Where("id = ? AND deleted_at IS NULL", userID).
		Scan(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewAPIError(constants.ErrUserNotFound, http.StatusNotFound)
		}
		return nil, errors.NewAPIError(constants.ErrInternalServerError, http.StatusInternalServerError)
	}

	userResponse := serializers.SerializeResponse(user, user.ResidentialDetails, user.OfficeDetails)
	response := serializers.SerializeGetUserResponse(user, userResponse)

	return &response, nil
}

func (s *UserService) ListUsers() ([]serializers.ListUserResponse, *errors.APIError) {
	var users []models.User

	if err := s.DB.Select("id, email").Find(&users).Error; err != nil {
		return nil, errors.NewAPIError(constants.ErrFailedToRetrieveUser, http.StatusInternalServerError)
	}

	response := serializers.SerializeListUser(users)

	return response, nil
}

func (s *UserService) UpdatePasswordByID(userID uint, req serializers.PasswordRequest) *errors.APIError {
	var user models.User
	if err := s.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return errors.NewAPIError(constants.ErrUserNotFound, http.StatusNotFound)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return errors.NewAPIError(constants.ErrInvalidOldPassword, http.StatusBadRequest)
	}

	newPasswordHash, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return errors.NewAPIError(constants.ErrFailedToHashPassword, http.StatusInternalServerError)
	}

	if err := s.DB.Model(&models.User{}).Where("id = ?", userID).Update("password", newPasswordHash).Error; err != nil {
		return errors.NewAPIError(constants.ErrFailedToUpdatePass, http.StatusInternalServerError)
	}

	return nil
}

func (s *UserService) DeleteUserByID(userID uint) (*serializers.DeleteUserResponse, *errors.APIError) {
	tx := s.DB.Begin()
	var user models.User

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Preload("OfficeDetails").Preload("ResidentialDetails").First(&user, "id = ?", userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			tx.Rollback()
			return nil, errors.NewAPIError(constants.ErrUserNotFound, http.StatusNotFound)
		}
		tx.Rollback()
		return nil, errors.NewAPIError(constants.ErrFailedToRetrieveUser, http.StatusInternalServerError)
	}

	if err := tx.Delete(&user).Error; err != nil {
		tx.Rollback()
		return nil, errors.NewAPIError(constants.ErrFailedToDeleteUser, http.StatusInternalServerError)
	}

	if err := tx.Delete(&user.ResidentialDetails).Error; err != nil {
		tx.Rollback()
		return nil, errors.NewAPIError(constants.ErrFailedToDeleteResAddr, http.StatusInternalServerError)
	}

	if err := tx.Delete(&user.OfficeDetails).Error; err != nil {
		tx.Rollback()
		return nil, errors.NewAPIError(constants.ErrFailedToDeleteOffAddr, http.StatusInternalServerError)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errors.NewAPIError(constants.ErrFailedToCommit, http.StatusInternalServerError)
	}

	userResponse := serializers.SerializeResponse(user, user.ResidentialDetails, user.OfficeDetails)
	response := serializers.SerializeDeleteUserResponse(user, userResponse)

	return &response, nil
}

func (s *UserService) UpdateUserByID(userID uint, req *serializers.UpdateUserRequest) (*serializers.GetUserResponse, *errors.APIError) {
	tx := s.DB.Begin()

	var user models.User
	if err := tx.Preload("OfficeDetails").Preload("ResidentialDetails").
		Where("id = ?", userID).First(&user).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewAPIError(constants.ErrUserNotFound, http.StatusNotFound)
		}
		return nil, errors.NewAPIError(constants.ErrFailedToRetrieveUser, http.StatusInternalServerError)
	}

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Gender != "" {
		var gender uint8
		var err *errors.APIError

		switch req.Gender {
		case "male":
			gender = 1
		case "female":
			gender = 2
		case "other":
			gender = 3
		default:
			err = errors.NewAPIError(constants.ErrInvalidGenderValue, http.StatusUnprocessableEntity)
		}

		if err != nil {
			tx.Rollback()
			return nil, err
		}
		user.Gender = gender
	}

	if req.DateOfBirth != "" {
		user.DateOfBirth = req.DateOfBirth
	}
	if req.MaritalStatus != "" {
		var maritalStatus uint8
		var err *errors.APIError

		switch req.MaritalStatus {
		case "single":
			maritalStatus = 1
		case "married":
			maritalStatus = 2
		default:
			err = errors.NewAPIError(constants.ErrInvalidMaritalStatus, http.StatusUnprocessableEntity)
		}

		if err != nil {
			tx.Rollback()
			return nil, err
		}
		user.MaritalStatus = maritalStatus
	}

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return nil, errors.NewAPIError(constants.ErrFailedToUpdate, http.StatusInternalServerError)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errors.NewAPIError(constants.ErrFailedToCommit, http.StatusInternalServerError)
	}

	userResponse := serializers.SerializeResponse(user, user.ResidentialDetails, user.OfficeDetails)
	response := serializers.SerializeGetUserResponse(user, userResponse)

	return &response, nil
}

func (s *UserService) checkIfUserExists(userID uint) bool {
	var user models.User
	err := s.DB.First(&user, userID).Error
	return err == nil
}

func (s *UserService) FollowUsersByID(userID uint, followerIDs []uint) *errors.APIError {
	tx := s.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, followerID := range followerIDs {

		if userID == followerID {
			tx.Rollback()
			return errors.NewAPIError(constants.ErrUserCantFollowItself, http.StatusBadRequest)
		}

		if !s.checkIfUserExists(followerID) {
			tx.Rollback()
			return errors.NewAPIError(fmt.Sprintf("user with ID %d does not exist", followerID), http.StatusBadRequest)
		}

		var follow models.UserFollowing
		err := tx.Where("user_id = ? AND follower_id = ? AND deleted_at IS NULL", userID, followerID).First(&follow).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {

				follow = models.UserFollowing{
					UserID:     userID,
					FollowerID: followerID,
				}

				if err := tx.Create(&follow).Error; err != nil {
					tx.Rollback()
					return errors.NewAPIError(constants.ErrFailedToFollow, http.StatusInternalServerError)
				}
			} else {
				tx.Rollback()
				return errors.NewAPIError(constants.ErrFailedToCheckFollowing, http.StatusInternalServerError)
			}
		} else {
			tx.Rollback()
			return errors.NewAPIError(fmt.Sprintf("User with ID %d is already following user with ID %d", userID, followerID), http.StatusBadRequest)
		}
	}

	if err := tx.Commit().Error; err != nil {
		return errors.NewAPIError(constants.ErrFailedToCommit, http.StatusInternalServerError)
	}

	return nil
}

func (s *UserService) UnfollowUsersByID(userID uint, followerIDs []uint) *errors.APIError {
	tx := s.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, followerID := range followerIDs {

		if !s.checkIfUserExists(followerID) {
			tx.Rollback()
			return errors.NewAPIError(fmt.Sprintf("User with ID %d does not exist", followerID), http.StatusBadRequest)
		}

		var userFollowing models.UserFollowing
		if err := tx.Where("user_id = ? AND follower_id = ? AND deleted_at IS NULL", userID, followerID).
			First(&userFollowing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				tx.Rollback()
				return errors.NewAPIError(fmt.Sprintf("User with ID %d is not following user with ID %d", userID, followerID), http.StatusBadRequest)
			}
			tx.Rollback()
			return errors.NewAPIError(constants.ErrFailedToUnfollow, http.StatusInternalServerError)
		}

		if err := tx.Where("user_id = ? AND follower_id = ? AND deleted_at IS NULL", userID, followerID).Delete(&models.UserFollowing{}).Error; err != nil {
			tx.Rollback()
			return errors.NewAPIError(constants.ErrFailedToUnfollow, http.StatusInternalServerError)
		}
	}

	if err := tx.Commit().Error; err != nil {
		return errors.NewAPIError(constants.ErrFailedToCommit, http.StatusInternalServerError)
	}

	return nil
}

func (s *UserService) GetFollowersByID(userID uint) ([]models.User, *errors.APIError) {
	var follower []models.User
	err := s.DB.Model(&models.User{}).
		Joins("JOIN user_followings ON user_followings.user_id = users.id").
		Where("user_followings.follower_id = ?", userID).
		Find(&follower).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewAPIError(constants.ErrUserNotFound, http.StatusNotFound)
		}
		return nil, errors.NewAPIError(err.Error(), http.StatusInternalServerError)
	}

	return follower, nil
}

func (s *UserService) GetFollowing(userID uint) ([]models.User, *errors.APIError) {
	var following []models.User
	err := s.DB.Model(&models.User{}).
		Joins("JOIN user_followings ON user_followings.follower_id = users.id").
		Where("user_followings.user_id = ?", userID).
		Find(&following).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewAPIError(constants.ErrUserNotFound, http.StatusNotFound)
		}
		return nil, errors.NewAPIError(err.Error(), http.StatusInternalServerError)
	}

	return following, nil
}
