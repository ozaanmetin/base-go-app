package repositories

import (
	"base-go-app/src/apps/regions/models"
	"errors"

	"gorm.io/gorm"
)

type CountryRepository struct {
	DB *gorm.DB
}

func CreateCountryRepository(db *gorm.DB) *CountryRepository {
	return &CountryRepository{DB: db}
}

func (repo *CountryRepository) FindAll() ([]models.Country, error) {
	var countries []models.Country
	if err := repo.DB.Find(&countries).Error; err != nil {
		return nil, err
	}
	return countries, nil
}

func (repo *CountryRepository) FindByID(id string) (*models.Country, error) {
	var country models.Country
	err := repo.DB.Where("id = ?", id).First(&country).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Country not found
			return nil, errors.New("country not found")
		}
		return nil, err
	}
	return &country, nil
}
