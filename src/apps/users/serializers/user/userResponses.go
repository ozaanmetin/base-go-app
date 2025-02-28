package serializers

type UserListResponse struct {
	ID       string
	Username string
	Role     string
}

type UserRetrieveResponse struct {
	ID        string 
	Username  string
	FirstName string
	LastName  string
	Email     string
	Role      string
}

