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

func LoginHandler(c *gin.Context) {
	var loginSerializer serializers.LoginRequest
	if err := c.ShouldBindJSON(&loginSerializer); err != nil {
		validationErrors := validations.GenericApiErrorValidator(err)
		c.JSON(
			http.StatusBadRequest,
			responses.ApiResponse{
				Message: messages.RequestValidationErrorMessage,
				Errors:  validationErrors,
			})
		return
	}

	userService := services.CreateUserService()
	// Attempt to login the user
	user, err := userService.Login(loginSerializer.Username, loginSerializer.Password)
	if err != nil {
		authErrors := []validations.ValidationError{
			{
				Message: err.Error(),
			},
		}

		c.JSON(
			http.StatusUnauthorized,
			responses.ApiResponse{
				Message: messages.InvalidCredentialsMessage,
				Errors:  authErrors,
			})
		return
	}
	// Generate JWT Token Pair
	accessToken, refreshToken, err := services.GenerateTokenPair(user)
	if err != nil {
		tokenGenerationErrors := []validations.ValidationError{
			{
				Message: err.Error(),
			},
		}
		c.JSON(http.StatusInternalServerError,
			responses.ApiResponse{
				Message: messages.TokenGenerationErrorMessage,
				Errors:  tokenGenerationErrors,
			})
		return
	}
	// Respond with the token pair
	response := responses.ApiResponse{
		Message: messages.LoginSuccessfulMessage,
		Data: serializers.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}}
	c.JSON(http.StatusOK, response)
}
