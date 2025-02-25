package serializers

type LoginRequest struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
}
