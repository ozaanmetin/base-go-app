package serializers

type RefreshRequest struct {
	RefreshToken string `binding:"required"`
}

type RefreshResponse struct {
	AccessToken string
}
