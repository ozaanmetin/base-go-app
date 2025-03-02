package handlers

import (
	messages "base-go-app/src/apps/users/handlers"
	serializers "base-go-app/src/apps/users/serializers/jwt"
	services "base-go-app/src/apps/users/services"
	responses "base-go-app/src/common/serializers/api"
	"base-go-app/src/common/serializers/errors"
	validations "base-go-app/src/common/utils/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RefreshHandler handles user token refresh requests
// @Summary Refresh user token
// @Description Refreshes a user token
// @Tags Users - Authentication | Public
// @Accept  json
// @Produce  json
// @Param refresh body serializers.RefreshRequest true "Refresh token"
// @Success 200 {object} responses.SuccessResponse{data=serializers.RefreshResponse}
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Router /api/auth/refresh [post]
func RefreshHandler(c *gin.Context) {
	var refreshSerializer serializers.RefreshRequest
	if err := c.ShouldBindJSON(&refreshSerializer); err != nil {
		validationErrors := validations.GenericApiErrorValidator(err)
		c.JSON(
			http.StatusBadRequest,
			responses.ErrorResponse{
				Message: messages.RequestValidationErrorMessage,
				Errors:  validationErrors,
			})
		return
	}

	// Attempt to refresh the token
	newAccessToken, err := services.RefreshToken(refreshSerializer.RefreshToken)
	if err != nil {
		tokenGenerationErrors := []errors.ErrorSerializer{
			{
				Message: err.Error(),
			},
		}
		c.JSON(http.StatusUnauthorized,
			responses.ErrorResponse{
				Message: messages.TokenGenerationErrorMessage,
				Errors:  tokenGenerationErrors,
			})
		return
	}
	response := responses.SuccessResponse{
		Message: messages.TokenRefreshedMessage,
		Data:    serializers.RefreshResponse{AccessToken: newAccessToken},
	}

	c.JSON(http.StatusOK, response)
}
