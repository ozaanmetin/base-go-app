package handlers

import (
	"base-go-app/src/apps/users/models"
	"base-go-app/src/apps/users/services"
	"net/http"

	serializers "base-go-app/src/apps/users/serializers/user"
	"base-go-app/src/common/pagination"
	responses "base-go-app/src/common/serializers/api"
	errors "base-go-app/src/common/serializers/errors"
	validations "base-go-app/src/common/utils/validators"

	"github.com/gin-gonic/gin"
)

type UserCrudHandler struct {
	UserService *services.UserService
}

func CreateUserCrudHandler() *UserCrudHandler {
	return &UserCrudHandler{UserService: services.CreateUserService()}
}

// ListAll godoc
// @Summary List all users
// @Description Get a paginated list of all users
// @Tags Users - User | Superuser - Manager
// @Produce json
// @Param page query int false "Page number"
// @Param pageSize query int false "Number of items per page"
// @Success 200 {object} responses.PaginatedResponse{data=serializers.UserListResponse}
// @Failure 500 {object} responses.ErrorResponse
// @Security BearerAuth
// @Router /api/users/ [get]
func (handler *UserCrudHandler) ListAll(c *gin.Context) {
	page, pageSize := pagination.GetPaginationParams(c)
	users, pagination, err := handler.UserService.FindAllPaginated(page, pageSize)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			responses.ErrorResponse{
				Message: "Error listing users",
				Errors:  []errors.ErrorSerializer{{Message: err.Error()}},
			},
		)
		return
	}

	var usersResponse []serializers.UserListResponse
	for _, user := range users {
		usersResponse = append(usersResponse, serializers.UserListResponse{
			ID:       user.ID.String(),
			Username: user.Username,
			Role:     user.Role,
		})
	}

	c.JSON(http.StatusOK, responses.PaginatedResponse{
		Message:    "Users listed successfully",
		Data:       usersResponse,
		Pagination: pagination,
	})
}

// Retrieve godoc
// @Summary Retrieve a user by ID
// @Description Get detailed information about a specific user
// @Tags Users - User | Superuser - Manager
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} responses.SuccessResponse{data=serializers.UserRetrieveResponse}
// @Failure 404 {object} responses.ErrorResponse
// @Security BearerAuth
// @Router /api/users/{id} [get]
func (handler *UserCrudHandler) Retrieve(c *gin.Context) {
	id := c.Param("id")
	user, err := handler.UserService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "User not found",
			Errors:  []errors.ErrorSerializer{{Message: err.Error()}},
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
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

// Create godoc
// @Summary Create a new user
// @Description Add a new user to the system
// @Tags Users - User | Superuser - Manager
// @Accept json
// @Produce json
// @Param user body serializers.UserCreateRequest true "User create request"
// @Success 201 {object} responses.SuccessResponse{data=serializers.UserRetrieveResponse}
// @Failure 400 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Security BearerAuth
// @Router /api/users/ [post]
func (handler *UserCrudHandler) Create(c *gin.Context) {
	var serializer serializers.UserCreateRequest
	if err := c.ShouldBindJSON(&serializer); err != nil {
		validationErrors := validations.GenericApiErrorValidator(err)
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
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
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Error creating user",
			Errors:  []errors.ErrorSerializer{{Message: err.Error()}},
		})
		return
	}

	c.JSON(http.StatusCreated, responses.SuccessResponse{
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

// Update godoc
// @Summary Update an existing user
// @Description Update user details by ID
// @Tags Users - User | Superuser - Manager
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body serializers.UserUpdateRequest true "User update request"
// @Success 200 {object} responses.SuccessResponse{data=serializers.UserRetrieveResponse}
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Security BearerAuth
// @Router /api/users/{id} [put]
func (handler *UserCrudHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var serializer serializers.UserUpdateRequest
	if err := c.ShouldBindJSON(&serializer); err != nil {
		validationErrors := validations.GenericApiErrorValidator(err)
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "Request validation error",
			Errors:  validationErrors,
		})
		return
	}

	user, err := handler.UserService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Message: "User not found",
			Errors:  []errors.ErrorSerializer{{Message: err.Error()}},
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
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Message: "Error updating user",
			Errors:  []errors.ErrorSerializer{{Message: err.Error()}},
		})
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse{
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

// Delete godoc
// @Summary Delete a user
// @Description Delete an existing user by ID
// @Tags Users - User | Superuser - Manager
// @Param id path string true "User ID"
// @Success 200 {object} responses.SuccessResponse
// @Failure 500 {object} responses.ErrorResponse
// @Security BearerAuth
// @Router /api/users/{id} [delete]
func (handler *UserCrudHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := handler.UserService.Delete(id)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			responses.ErrorResponse{
				Message: "Error deleting user",
				Errors: []errors.ErrorSerializer{
					{
						Message: err.Error(),
					},
				},
			})
		return
	}

	c.JSON(
		http.StatusOK,
		responses.SuccessResponse{
			Message: "User deleted successfully",
		})
}
