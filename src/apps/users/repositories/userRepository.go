package repositories

import (
	"base-go-app/src/apps/users/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func CreateUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	if err := repo.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	err := repo.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// User not found
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := repo.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// User not found
			return nil, errors.New("user not found")
		}
		return nil, err

	}
	return &user, nil
}

func (repo *UserRepository) Create(user *models.User) error {
	return repo.DB.Create(user).Error
}

func (repo *UserRepository) Update(user *models.User) error {
	return repo.DB.Save(user).Error
}

func (repo *UserRepository) Delete(id string) error {
	var user models.User
	if err := repo.DB.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("user not found")
		}
		return err
	}

	if err := repo.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
