package services

import (
	"mini-social-network/errors"
	"mini-social-network/mocks"
	"mini-social-network/serializers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserServiceInt(t *testing.T) {

	mockService := mocks.NewUserServiceInt(t)

	userID := uint(1)
	followerIDs := []uint{2, 3, 4}

	t.Run("Test GetUserByID - Success", func(t *testing.T) {
		expectedResponse := &serializers.GetUserResponse{}

		mockService.On("GetUserByID", userID).Return(expectedResponse, nil)

		response, err := mockService.GetUserByID(userID)

		assert.Nil(t, err)
		assert.Equal(t, expectedResponse, response)

		mockService.AssertExpectations(t)
	})

	t.Run("Test GetUserByID - Error", func(t *testing.T) {
		mockService.On("GetUserByID", userID).Return(nil, errors.NewAPIError("User not found", 404))

		response, err := mockService.GetUserByID(userID)

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.Equal(t, "User not found", err.Error())

		mockService.AssertExpectations(t)
	})

	t.Run("Test ListUsers - Success", func(t *testing.T) {
		expectedResponse := []serializers.ListUserResponse{
			{UserId: 1, Email: "test@example.com"},
		}

		mockService.On("ListUsers").Return(expectedResponse, nil)

		response, err := mockService.ListUsers()

		assert.Nil(t, err)
		assert.Equal(t, expectedResponse, response)

		mockService.AssertExpectations(t)
	})

	t.Run("Test ListUsers - Error", func(t *testing.T) {
		mockService.On("ListUsers").Return(nil, errors.NewAPIError("Failed to retrieve users", 500))

		response, err := mockService.ListUsers()

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.Equal(t, "Failed to retrieve users", err.Message)

		mockService.AssertExpectations(t)
	})

	t.Run("Test FollowUsersByID - Success", func(t *testing.T) {
		mockService.On("FollowUsersByID", userID, followerIDs).Return(nil)

		err := mockService.FollowUsersByID(userID, followerIDs)

		assert.Nil(t, err)

		mockService.AssertExpectations(t)
	})

	t.Run("Test FollowUsersByID - Error", func(t *testing.T) {
		mockService.On("FollowUsersByID", userID, followerIDs).Return(errors.NewAPIError("Cannot follow users", 400))

		err := mockService.FollowUsersByID(userID, followerIDs)

		assert.NotNil(t, err)
		assert.Equal(t, "Cannot follow users", err.Message)

		mockService.AssertExpectations(t)
	})

	t.Run("Test DeleteUserByID - Success", func(t *testing.T) {
		expectedResponse := &serializers.DeleteUserResponse{}

		mockService.On("DeleteUserByID", userID).Return(expectedResponse, nil)

		response, err := mockService.DeleteUserByID(userID)

		assert.Nil(t, err)
		assert.Equal(t, expectedResponse, response)

		mockService.AssertExpectations(t)
	})

	t.Run("Test DeleteUserByID - Error", func(t *testing.T) {
		mockService.On("DeleteUserByID", userID).Return(nil, errors.NewAPIError("Failed to delete user", 500))

		response, err := mockService.DeleteUserByID(userID)

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.Equal(t, "Failed to delete user", err.Message)

		mockService.AssertExpectations(t)
	})

	t.Run("Test UpdateUserByID - Success", func(t *testing.T) {
		req := &serializers.UpdateUserRequest{}

		expectedResponse := &serializers.GetUserResponse{}

		mockService.On("UpdateUserByID", userID, req).Return(expectedResponse, nil)

		response, err := mockService.UpdateUserByID(userID, req)

		assert.Nil(t, err)
		assert.Equal(t, expectedResponse, response)

		mockService.AssertExpectations(t)
	})

	t.Run("Test UpdateUserByID - Error", func(t *testing.T) {
		req := &serializers.UpdateUserRequest{}

		mockService.On("UpdateUserByID", userID, req).Return(nil, errors.NewAPIError("Failed to update user", 500))

		response, err := mockService.UpdateUserByID(userID, req)

		assert.Nil(t, response)
		assert.NotNil(t, err)
		assert.Equal(t, "Failed to update user", err.Message)

		mockService.AssertExpectations(t)
	})

	t.Run("Test UnfollowUsersByID - Success", func(t *testing.T) {
		mockService.On("UnfollowUsersByID", userID, followerIDs).Return(nil)

		err := mockService.UnfollowUsersByID(userID, followerIDs)

		assert.Nil(t, err)

		mockService.AssertExpectations(t)
	})

	t.Run("Test UnfollowUsersByID - Error", func(t *testing.T) {
		mockService.On("UnfollowUsersByID", userID, followerIDs).Return(errors.NewAPIError("Failed to unfollow users", 500))

		err := mockService.UnfollowUsersByID(userID, followerIDs)

		assert.NotNil(t, err)
		assert.Equal(t, "Failed to unfollow users", err.Message)

		mockService.AssertExpectations(t)
	})
}
