package middlewares

import (
	"base-go-app/src/common/utils/environment"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	corsAllowAllOrigins := environment.GetAsBool("CORS_ALLOW_ALL_ORIGINS", false)
	allowedOrigins := environment.GetAsSlice("CORS_ALLOWED_ORIGINS", []string{""})
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		if corsAllowAllOrigins {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		} else {
			// Loop through allowed origins and check if the request origin is allowed
			isAllowed := false
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					isAllowed = true
					break
				}
			}
			// If the origin is allowed, set the Access-Control-Allow-Origin header
			// Otherwise return a 403 Forbidden response
			if isAllowed {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			} else {
				c.AbortWithStatusJSON(
					http.StatusForbidden,
					gin.H{"error": "CORS not allowed for this origin"},
				)
				return
			}
		}
		// Set the other CORS headers
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}
}
