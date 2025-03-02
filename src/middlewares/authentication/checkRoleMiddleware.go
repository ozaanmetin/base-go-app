package middlewares

import (
	"net/http"

	responses "base-go-app/src/common/serializers/api"
	"base-go-app/src/common/serializers/errors"

	"github.com/gin-gonic/gin"
)

// CheckRoleMiddleware to enforce role-based access control
func CheckRoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				responses.ErrorResponse{
					Message: "Unauthorized Access",
					Errors: []errors.ErrorSerializer{{
						Field:   "Authorization",
						Message: "No role found in token",
					}},
				},
			)
			return
		}

		// Check if the user's role is one of the allowed roles
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				c.Next()
				return
			}
		}

		// If the role doesn't match, abort with an error
		c.AbortWithStatusJSON(
			http.StatusForbidden,
			responses.ErrorResponse{
				Message: "Access Denied",
				Errors: []errors.ErrorSerializer{{
					Field:   "Authorization",
					Message: "No permission to access this resource",
				}},
			},
		)
	}
}
