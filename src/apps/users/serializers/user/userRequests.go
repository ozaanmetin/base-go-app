package serializers

type UserCreateRequest struct {
	Username  string `binding:"required"`
	FirstName string `binding:"required"`
	LastName  string `binding:"required"`
	Email     string `binding:"required"`
	Password  string `binding:"required"`
	Role      string
}

type UserUpdateRequest struct {
	Username  *string
	FirstName *string
	LastName  *string
	Email     *string
}
