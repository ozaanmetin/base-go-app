package handlers

import (
	"base-go-app/src/apps/regions/services"
	"base-go-app/src/common/pagination"
	responses "base-go-app/src/common/serializers/api"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CountryHandler struct {
	CountryService *services.CountryService
}

func CreateCountryHandler() *CountryHandler {
	return &CountryHandler{CountryService: services.CreateCountryService()}
}

func (handler *CountryHandler) ListAll(c *gin.Context) {
	page, pageSize := pagination.GetPaginationParams(c)
	countries, pagination, err := handler.CountryService.FindAllPaginated(page, pageSize)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			responses.ApiResponse{
				Message: "Error listing users",
				Errors: []responses.ApiResponse{
					{
						Message: err.Error(),
					},
				},
			})
		return
	}
	c.JSON(
		http.StatusOK,
		responses.ApiResponse{
			Message:    "Countries listed successfully",
			Data:       countries,
			Pagination: pagination,
		})
}

func (handler *CountryHandler) Retrieve(c *gin.Context) {
	id := c.Param("id")
	// print type
	fmt.Printf("Type of id: %T\n", id)
	country, err := handler.CountryService.FindByID(id)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			responses.ApiResponse{
				Message: "Error retrieving country",
				Errors: []responses.ApiResponse{
					{
						Message: err.Error(),
					},
				},
			})
		return
	}
	c.JSON(
		http.StatusOK,
		responses.ApiResponse{
			Message: "Country retrieved successfully",
			Data:    country,
		})

}
