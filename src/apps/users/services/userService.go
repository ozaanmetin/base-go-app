package services

import (
	"base-go-app/src/apps/users/models"
	"base-go-app/src/apps/users/repositories"
	"base-go-app/src/database"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func CreateUserService() *UserService {
	return &UserService{UserRepository: repositories.CreateUserRepository(database.PostgresContext)}
}

func (service *UserService) Login(username string, password string) (*models.User, error) {
	user, err := service.UserRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	err = user.ComparePassword(password)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (service *UserService) FindAll() ([]models.User, error) {
	return service.UserRepository.FindAll()
}

func (service *UserService) FindByID(id string) (*models.User, error) {
	return service.UserRepository.FindByID(id)
}

func (service *UserService) Create(user *models.User) error {
	return service.UserRepository.Create(user)
}

func (service *UserService) Update(user *models.User) error {
	return service.UserRepository.Update(user)
}

func (service *UserService) Delete(id string) error {
	return service.UserRepository.Delete(id)
}
