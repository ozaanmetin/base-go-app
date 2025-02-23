package models

import (
	BaseModels "base-go-app/src/common/models/essentials"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseModels.BaseUUIDModel
	BaseModels.BaseTemporalModel
	BaseModels.BaseActivenessModel
	Username  string `gorm:"unique;not null;size:50" json:"username"`
	FirstName string `gorm:"not null;size:100" json:"first_name"`
	LastName  string `gorm:"not null;size:100" json:"last_name"`
	Email     string `gorm:"unique;not null;size:255" json:"email"`
	Password  string `gorm:"not null;size:255" json:"password"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Deactivate(db *gorm.DB) error {
	u.IsActive = false
	return db.Save(u).Error
}

func (u *User) Activate(db *gorm.DB) error {
	u.IsActive = true
	return db.Save(u).Error
}
