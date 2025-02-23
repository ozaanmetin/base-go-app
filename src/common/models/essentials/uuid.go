package models

import (
	"github.com/google/uuid"
)

type BaseUUIDModel struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v7()" json:"id"`
}
