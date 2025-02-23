package responses

type ApiResponse struct {
	Message string
	Data    interface{}
	Errors  interface{}
}
