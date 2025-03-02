package serializers

type UserCreateRequest struct {
	Username  string `binding:"required" json:"username"`
	FirstName string `binding:"required" json:"firstName"`
	LastName  string `binding:"required" json:"lastName"`
	Email     string `binding:"required" json:"email"`
	Password  string `binding:"required" json:"password"`
	Role      string `json:"role"`
}

type UserUpdateRequest struct {
	Username  *string `json:"username"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     *string `json:"email"`
}
