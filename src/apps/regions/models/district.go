package models

import (
	BaseModels "base-go-app/src/common/models/essentials"

	"github.com/google/uuid"
)

type District struct {
	BaseModels.BaseUUIDModel
	BaseModels.BaseTemporalModel
	Name      string     `gorm:"not null;size:255" json:"name"`
	CountryID uuid.UUID  `gorm:"type:uuid;not null" json:"country_id"`
	RegionID  *uuid.UUID `gorm:"type:uuid;null" json:"region_id"`
	CityID    uuid.UUID  `gorm:"type:uuid;not null" json:"city_id"`

	// Relationships
	Country Country `gorm:"foreignKey:CountryID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"country"`
	Region  *Region `gorm:"foreignKey:RegionID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"region"`
	City    City    `gorm:"foreignKey:CityID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"city"`
}

func (District) TableName() string {
	return "districts"
}
