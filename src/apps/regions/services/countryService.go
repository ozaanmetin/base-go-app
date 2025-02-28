package services

import (
	"base-go-app/src/apps/regions/models"
	"base-go-app/src/apps/regions/repositories"
	"base-go-app/src/common/pagination"

	"base-go-app/src/database"
)

type CountryService struct {
	CountryRepository *repositories.CountryRepository
}

func CreateCountryService() *CountryService {
	return &CountryService{CountryRepository: repositories.CreateCountryRepository(database.PostgresContext)}
}

func (service *CountryService) FindAll() ([]models.Country, error) {
	return service.CountryRepository.FindAll()
}

func (service *CountryService) FindAllPaginated(page int, pageSize int) ([]models.Country, *pagination.Pagination, error) {
	var countries []models.Country
	pagination, err := pagination.Paginate(service.CountryRepository.DB, &models.Country{}, page, pageSize, &countries)
	if err != nil {
		return nil, nil, err
	}
	return countries, pagination, err
}

func (service *CountryService) FindByID(id string) (*models.Country, error) {
	return service.CountryRepository.FindByID(id)
}
