package models

import (
	"gorm.io/gorm"
)

type BaseSoftDeleteModel struct {
	IsDeleted bool           `gorm:"default:false" json:"is_deleted"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
