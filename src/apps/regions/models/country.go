package models

import (
	BaseModels "base-go-app/src/common/models/essentials"
	"base-go-app/src/common/utils/types"
)

type Country struct {
	BaseModels.BaseUUIDModel
	BaseModels.BaseTemporalModel
	// JSON is used for name in order to keep the name of countries for different languages
	Name string `gorm:"type:jsonb;not null" json:"name"`
	Code string `gorm:"unique;not null;size:10" json:"code"`
}

func (Country) TableName() string {
	return "countries"
}

func (c *Country) GetNameAsMap() (map[string]interface{}, error) {
	nameMap, err := types.GetJsonAsMap(c.Name)
	if err != nil {
		return nil, err
	}
	return nameMap, nil
}
