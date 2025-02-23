package models

type BaseActivenessModel struct {
	IsActive bool `gorm:"default:true" json:"is_active"`
}
