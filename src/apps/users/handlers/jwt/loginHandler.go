package handlers

import (
	messages "base-go-app/src/apps/users/handlers"
	serializers "base-go-app/src/apps/users/serializers/jwt"
	services "base-go-app/src/apps/users/services"
	responses "base-go-app/src/common/serializers/api"
	errors "base-go-app/src/common/serializers/errors"
	validations "base-go-app/src/common/utils/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginHandler handles user login requests
// @Summary User login
// @Description Authenticates a user and returns an access token and refresh token
// @Tags Users - Authentication | Public
// @Accept  json
// @Produce  json
// @Param login body serializers.LoginRequest true "User login credentials"
// @Success 200 {object} responses.SuccessResponse{data=serializers.LoginResponse}
// @Failure 400 {object} responses.ErrorResponse
// @Failure 401 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/auth/login [post]
func LoginHandler(c *gin.Context) {
	var loginSerializer serializers.LoginRequest
	if err := c.ShouldBindJSON(&loginSerializer); err != nil {
		validationErrors := validations.GenericApiErrorValidator(err)
		c.JSON(
			http.StatusBadRequest,
			responses.ErrorResponse{
				Message: messages.RequestValidationErrorMessage,
				Errors:  validationErrors,
			})
		return
	}

	userService := services.CreateUserService()
	// Attempt to login the user
	user, err := userService.Login(loginSerializer.Username, loginSerializer.Password)
	if err != nil {
		authErrors := []errors.ErrorSerializer{
			{
				Message: err.Error(),
			},
		}

		c.JSON(
			http.StatusUnauthorized,
			responses.ErrorResponse{
				Message: messages.InvalidCredentialsMessage,
				Errors:  authErrors,
			})
		return
	}
	// Generate JWT Token Pair
	accessToken, refreshToken, err := services.GenerateTokenPair(user)
	if err != nil {
		tokenGenerationErrors := []errors.ErrorSerializer{
			{
				Message: err.Error(),
			},
		}
		c.JSON(http.StatusInternalServerError,
			responses.ErrorResponse{
				Message: messages.TokenGenerationErrorMessage,
				Errors:  tokenGenerationErrors,
			})
		return
	}
	// Respond with the token pair
	response := responses.SuccessResponse{
		Message: messages.LoginSuccessfulMessage,
		Data: serializers.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}}
	c.JSON(http.StatusOK, response)
}
