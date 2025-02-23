package handlers

import (
	messages "base-go-app/src/apps/users/handlers"
	serializers "base-go-app/src/apps/users/serializers/jwt"
	services "base-go-app/src/apps/users/services"
	responses "base-go-app/src/common/serializers/api"
	validations "base-go-app/src/common/utils/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RefreshHandler(c *gin.Context) {
	var refreshSerializer serializers.RefreshRequest
	if err := c.ShouldBindJSON(&refreshSerializer); err != nil {
		validationErrors := validations.GenericApiErrorValidator(err)
		c.JSON(
			http.StatusBadRequest,
			responses.ApiResponse{
				Message: messages.RequestValidationErrorMessage,
				Errors:  validationErrors,
			})
		return
	}

	// Attempt to refresh the token
	newAccessToken, err := services.RefreshToken(refreshSerializer.RefreshToken)
	if err != nil {
		tokenGenerationErrors := []validations.ValidationError{
			{
				Message: err.Error(),
			},
		}
		c.JSON(http.StatusUnauthorized,
			responses.ApiResponse{
				Message: messages.TokenGenerationErrorMessage,
				Errors:  tokenGenerationErrors,
			})
		return
	}
	response := responses.ApiResponse{
		Message: messages.TokenRefreshedMessage,
		Data:    serializers.RefreshResponse{AccessToken: newAccessToken},
	}

	c.JSON(http.StatusOK, response)
}
