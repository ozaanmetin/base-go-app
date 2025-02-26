package handlers

import (
	"base-go-app/src/apps/users/models"
	"base-go-app/src/apps/users/services"
	"net/http"

	serializers "base-go-app/src/apps/users/serializers/user"
	"base-go-app/src/common/pagination"
	responses "base-go-app/src/common/serializers/api"
	validations "base-go-app/src/common/utils/validators"

	"github.com/gin-gonic/gin"
)

type UserCrudHandler struct {
	UserService *services.UserService
}

func CreateUserCrudHandler() *UserCrudHandler {
	return &UserCrudHandler{UserService: services.CreateUserService()}
}

func (handler *UserCrudHandler) ListAll(c *gin.Context) {
	page, pageSize := pagination.GetPaginationParams(c)
	users, pagination, err := handler.UserService.FindAllPaginated(page, pageSize)
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
	}
	c.JSON(
		http.StatusOK,
		responses.ApiResponse{
			Message:    "Users listed successfully",
			Data:       usersResponse,
			Pagination: pagination,
		})
}

func (handler *UserCrudHandler) Retrieve(c *gin.Context) {
	id := c.Param("id")
	user, err := handler.UserService.FindByID(id)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			responses.ApiResponse{
				Message: "User not found",
				Errors: []responses.ApiResponse{
					{
						Message: err.Error(),
					},
				},
			})
		return
	}

	c.JSON(
		http.StatusOK,
		responses.ApiResponse{
			Message: "User retrieved successfully",
			Data: serializers.UserRetrieveResponse{
				ID:        user.ID.String(),
				Username:  user.Username,
				Role:      user.Role,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
			},
		})
}

func (handler *UserCrudHandler) Create(c *gin.Context) {
	var serializer serializers.UserCreateRequest
	if err := c.ShouldBindJSON(&serializer); err != nil {
		validationErrors := validations.GenericApiErrorValidator(err)
		c.JSON(
			http.StatusBadRequest,
			responses.ApiResponse{
				Message: "Request validation error",
				Errors:  validationErrors,
			})
		return
	}

	user := models.User{
		Username:  serializer.Username,
		FirstName: serializer.FirstName,
		LastName:  serializer.LastName,
		Email:     serializer.Email,
		Role:      serializer.Role,
	}

	err := handler.UserService.Create(&user)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			responses.ApiResponse{
				Message: "Error creating user",
				Errors: []responses.ApiResponse{
					{
						Message: err.Error(),
					},
				},
			})
		return
	}

	c.JSON(
		http.StatusCreated,
		responses.ApiResponse{
			Message: "User created successfully",
			Data: serializers.UserRetrieveResponse{
				ID:        user.ID.String(),
				Username:  user.Username,
				Role:      user.Role,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
			},
		})

}

func (handler *UserCrudHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var serializer serializers.UserUpdateRequest
	if err := c.ShouldBindJSON(&serializer); err != nil {
		validationErrors := validations.GenericApiErrorValidator(err)
		c.JSON(
			http.StatusBadRequest,
			responses.ApiResponse{
				Message: "Request validation error",
				Errors:  validationErrors,
			})
		return
	}

	user, err := handler.UserService.FindByID(id)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			responses.ApiResponse{
				Message: "User not found",
				Errors: []responses.ApiResponse{
					{
						Message: err.Error(),
					},
				},
			})
		return
	}

	if serializer.Username != nil {
		user.Username = *serializer.Username
	}
	if serializer.FirstName != nil {
		user.FirstName = *serializer.FirstName
	}
	if serializer.LastName != nil {
		user.LastName = *serializer.LastName
	}
	if serializer.Email != nil {
		user.Email = *serializer.Email
	}

	err = handler.UserService.Update(user)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			responses.ApiResponse{
				Message: "Error updating user",
				Errors: []responses.ApiResponse{
					{
						Message: err.Error(),
					},
				},
			})
		return
	}

	c.JSON(
		http.StatusOK,
		responses.ApiResponse{
			Message: "User updated successfully",
			Data: serializers.UserRetrieveResponse{
				ID:        user.ID.String(),
				Username:  user.Username,
				Role:      user.Role,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
			},
		})
}

func (handler *UserCrudHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := handler.UserService.Delete(id)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			responses.ApiResponse{
				Message: "Error deleting user",
				Errors: []responses.ApiResponse{
					{
						Message: err.Error(),
					},
				},
			})
		return
	}

	c.JSON(
		http.StatusOK,
		responses.ApiResponse{
			Message: "User deleted successfully",
		})
}
