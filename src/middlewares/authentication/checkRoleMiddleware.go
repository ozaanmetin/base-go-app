package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckRoleMiddleware to enforce role-based access control
func CheckRoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Role not found in token"})
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
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Not allowed to access this resource"})
	}
}
