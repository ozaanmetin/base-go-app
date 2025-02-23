package models

import (
	"time"
)

type BaseTemporalModel struct {
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;default:null" json:"updated_at"`
}
