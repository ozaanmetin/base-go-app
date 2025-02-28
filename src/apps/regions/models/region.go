package models

import (
	BaseModels "base-go-app/src/common/models/essentials"

	"github.com/google/uuid"
)

type Region struct {
	BaseModels.BaseUUIDModel
	BaseModels.BaseTemporalModel
	Name      string     `gorm:"not null;size:255" json:"name"`
	CountryID uuid.UUID  `gorm:"type:uuid;not null" json:"country_id"`
	ParentID  *uuid.UUID `gorm:"type:uuid;null" json:"parent_id"`

	// Relationships
	Country  Country  `gorm:"foreignKey:CountryID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"country"`
	Parent   *Region  `gorm:"foreignKey:ParentID;constraint:OnDelete:SET NULL,OnUpdate:CASCADE;" json:"parent"`
	Children []Region `gorm:"foreignKey:ParentID" json:"children"`
}

func (Region) TableName() string {
	return "regions"
}
