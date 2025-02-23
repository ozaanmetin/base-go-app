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

func (repo *UserRepository) FindAll() {
	// Implementation here
}

func (repo *UserRepository) FindByID() {
	// Implementation here
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
