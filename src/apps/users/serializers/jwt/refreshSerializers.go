package serializers

type RefreshRequest struct {
	RefreshToken string `binding:"required" json:"refreshToken"`
}

type RefreshResponse struct {
	AccessToken string `json:"accessToken"`
}
