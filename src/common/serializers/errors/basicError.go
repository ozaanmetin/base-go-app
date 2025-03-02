package errors

type ErrorSerializer struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
