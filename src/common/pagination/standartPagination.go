package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Pagination struct
type Pagination struct {
	TotalRecords int `json:"totalRecords"`
	TotalPages   int `json:"totalPages"`
	Page         int `json:"page"`
	PageSize     int `json:"pageSize"`
}

func GetPaginationParams(c *gin.Context) (int, int) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	// Max page size is 200
	// TODO: Add this to environment variables
	if err != nil || pageSize > 200 {
		pageSize = 10
	}
	return page, pageSize
}

func Paginate(db *gorm.DB, model interface{}, page int, pageSize int, resultData interface{}) (*Pagination, error) {
	// Get total records
	var totalRecords int64
	db.Model(model).Count(&totalRecords)
	// Get data
	offset := (page - 1) * pageSize
	query := db.Offset(offset).Limit(pageSize).Find(resultData)
	if query.Error != nil {
		return nil, query.Error
	}
	// Calculate total pages
	totalPages := int((totalRecords + int64(pageSize) - 1) / int64(pageSize)) // Round up
	// Create pagination object
	pagination := &Pagination{
		TotalRecords: int(totalRecords),
		TotalPages:   totalPages,
		Page:         page,
		PageSize:     pageSize,
	}
	return pagination, nil
}
