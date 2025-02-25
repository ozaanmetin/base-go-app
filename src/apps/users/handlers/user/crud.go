package handlers

import (
	"base-go-app/src/apps/users/services"
	"net/http"

	serializers "base-go-app/src/apps/users/serializers/user"
	responses "base-go-app/src/common/serializers/api"

	"github.com/gin-gonic/gin"
)

type UserCrudHandler struct {
	UserService *services.UserService
}

func CreateUserCrudHandler() *UserCrudHandler {
	return &UserCrudHandler{UserService: services.CreateUserService()}
}

func (handler *UserCrudHandler) ListAll(c *gin.Context) {
	users, err := handler.UserService.FindAll()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			responses.ApiResponse{
				Message: "Error listing users",
				Errors: []responses.ApiResponse{
					{
						Message: err.Error(),
					},
				},
			})
		return
	}

	var usersResponse []serializers.UserListResponse

	for _, user := range users {
		usersResponse = append(
			usersResponse,
			serializers.UserListResponse{
				ID:       user.ID.String(),
				Username: user.Username,
				Role:     user.Role,
			})

	c.JSON(
		http.StatusOK,
		responses.ApiResponse{
			Message: "Users listed successfully",
			Data:    usersResponse,
		})

	}
}
