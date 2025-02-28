package models

import (
	"base-go-app/src/common/fields"
	BaseModels "base-go-app/src/common/models/essentials"
)

type Country struct {
	BaseModels.BaseUUIDModel
	BaseModels.BaseTemporalModel
	// JSON is used for name in order to keep the name of countries for different languages
	Name fields.Jsonb `gorm:"type:jsonb;not null" json:"name"`
	Code string       `gorm:"unique;not null;size:10" json:"code"`
}

func (Country) TableName() string {
	return "countries"
}
