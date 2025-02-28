package services

import (
	"base-go-app/config/settings/environment"
	"base-go-app/src/apps/users/models"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateTokenPair generates an access token and a refresh token for a user
func GenerateTokenPair(user *models.User) (string, string, error) {
	accessTokenLifetime := time.Duration(environment.GetAsInt("ACCESS_TOKEN_LIFETIME", 5)) * time.Minute
	refreshTokenLifetime := time.Duration(environment.GetAsInt("REFRESH_TOKEN_LIFETIME", 30)) * 24 * time.Hour
	// Creating the access token
	accessTokenClaims := jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(accessTokenLifetime).Unix(),
		"iat":  time.Now().Unix(),
	}
	accessTokenString, err := createToken(accessTokenClaims)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate access token: %w", err)
	}
	// Creating the refresh token
	refreshTokenClaims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(refreshTokenLifetime).Unix(),
		"iat": time.Now().Unix(),
	}
	refreshTokenString, err := createToken(refreshTokenClaims)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return accessTokenString, refreshTokenString, nil
}

// RefreshToken generates a new access token from a refresh token
func RefreshToken(refreshToken string) (string, error) {
	signingKey := []byte(os.Getenv("JWT_SIGNING_KEY"))
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid or expired refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["sub"] == nil {
		return "", fmt.Errorf("invalid token claims")
	}

	userID := claims["sub"]
	accessTokenLifetime := time.Duration(environment.GetAsInt("ACCESS_TOKEN_LIFETIME", 5)) * time.Minute
	// Generate a new access token
	newAccessTokenClaims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(accessTokenLifetime).Unix(),
		"iat": time.Now().Unix(),
	}
	// Create the new access token
	newAccessToken, err := createToken(newAccessTokenClaims)
	if err != nil {
		return "", fmt.Errorf("failed to generate new access token: %w", err)
	}
	return newAccessToken, nil
}

// createToken generates a JWT token with the provided claims
func createToken(claims jwt.Claims) (string, error) {
	signingKey := []byte(os.Getenv("JWT_SIGNING_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return tokenString, nil
}
