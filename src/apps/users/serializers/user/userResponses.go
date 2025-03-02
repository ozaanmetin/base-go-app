package serializers

type UserListResponse struct {
	ID       string 	`json:"id"`
	Username string 	`json:"username"`
	Role     string 	`json:"role"`
}

type UserRetrieveResponse struct {
	ID        string	`json:"id"`
	Username  string	`json:"username"`
	FirstName string	`json:"firstName"`
	LastName  string	`json:"lastName"`
	Email     string	`json:"email"`
	Role      string	`json:"role"`
}
