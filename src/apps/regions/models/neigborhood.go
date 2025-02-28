package models

import (
	BaseModels "base-go-app/src/common/models/essentials"

	"github.com/google/uuid"
)

type Neighborhood struct {
	BaseModels.BaseUUIDModel
	BaseModels.BaseTemporalModel
	Name       string     `gorm:"not null;size:255" json:"name"`
	CountryID  uuid.UUID  `gorm:"type:uuid;not null" json:"country_id"`
	RegionID   *uuid.UUID `gorm:"type:uuid;null" json:"region_id"`
	CityID     uuid.UUID  `gorm:"type:uuid;not null" json:"city_id"`
	DistrictID uuid.UUID  `gorm:"type:uuid;not null" json:"district_id"`

	// Relationships
	Country  Country  `gorm:"foreignKey:CountryID;constraint:OnDelete:CASCADE;" json:"country"`
	Region   *Region  `gorm:"foreignKey:RegionID;constraint:OnDelete:SET NULL;" json:"region"`
	City     City     `gorm:"foreignKey:CityID;constraint:OnDelete:CASCADE;" json:"city"`
	District District `gorm:"foreignKey:DistrictID;constraint:OnDelete:CASCADE;" json:"district"`
}

func (Neighborhood) TableName() string {
	return "neighborhoods"
}
