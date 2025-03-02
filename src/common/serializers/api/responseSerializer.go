package responses

import (
	"base-go-app/src/common/pagination"
	"base-go-app/src/common/serializers/errors"
)

type PaginatedResponse struct {
	Message    string                 `json:"message"`
	Data       interface{}            `json:"data"`
	Pagination *pagination.Pagination `json:"pagination"`
}

type ErrorResponse struct {
	Message string                   `json:"message"`
	Errors  []errors.ErrorSerializer `json:"errors"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
