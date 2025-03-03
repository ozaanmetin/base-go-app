package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	responses "base-go-app/src/common/serializers/api"
	"base-go-app/src/common/serializers/errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	jwtSigningKey := []byte(os.Getenv("JWT_SIGNING_KEY"))
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				responses.ErrorResponse{
					Message: "Authorization header is required",
					Errors: []errors.ErrorSerializer{{
						Field:   "Authorization",
						Message: "Authorization header is missing or empty",
					}},
				},
			)
			return
		}
		// Trim the Bearer prefix and validate the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSigningKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				responses.ErrorResponse{
					Message: "Unauthorized Access",
					Errors: []errors.ErrorSerializer{{
						Field:   "Authorization",
						Message: "Invalid Token",
					}},
				},
			)
			return
		}
		// Set the granted userId in the context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId", claims["sub"])
			c.Set("role", claims["role"])
			c.Next()
		} else {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				responses.ErrorResponse{
					Message: "Unauthorized Access",
					Errors: []errors.ErrorSerializer{{
						Field:   "Authorization",
						Message: "Invalid Token",
					}},
				},
			)
		}
	}
}
