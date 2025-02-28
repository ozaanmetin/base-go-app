package models

import (
	BaseModels "base-go-app/src/common/models/essentials"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	Superuser = "superuser"
	Manager   = "manager"
	Member    = "member"
)

var validRoles = map[string]bool{
	Superuser: true,
	Manager:   true,
	Member:    true,
}

type User struct {
	BaseModels.BaseUUIDModel
	BaseModels.BaseTemporalModel
	BaseModels.BaseActivenessModel
	Username  string `gorm:"unique;not null;size:50" json:"username"`
	FirstName string `gorm:"not null;size:100" json:"first_name"`
	LastName  string `gorm:"not null;size:100" json:"last_name"`
	Email     string `gorm:"unique;not null;size:255" json:"email"`
	Password  string `gorm:"not null;size:255" json:"password"`
	Role      string `gorm:"not null;size:50;default:'member'" json:"role"`
}

func (User) TableName() string {
	return "users"
}

// BeforeCreate hashes the password and validates the role before the record is created
func (u *User) BeforeCreate(db *gorm.DB) error {
	// Hash password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Validate role
	if !u.isValidRole() {
		return errors.New("invalid role: " + u.Role)
	}

	// Set hashed password
	u.Password = string(hashedPassword)
	return nil
}

// isValidRole checks if the user's role is valid
func (u *User) isValidRole() bool {
	if u.Role != "" {
		_, exists := validRoles[u.Role]
		return exists
	}
	return true

}

// Deactivate marks the user as inactive
func (u *User) Deactivate(db *gorm.DB) error {
	u.IsActive = false
	return db.Save(u).Error
}

// Activate marks the user as active
func (u *User) Activate(db *gorm.DB) error {
	u.IsActive = true
	return db.Save(u).Error
}

// Role checkers
func (u *User) IsSuperuser() bool {
	return u.Role == Superuser
}

func (u *User) IsManager() bool {
	return u.Role == Manager
}

func (u *User) IsMember() bool {
	return u.Role == Member
}

// ComparePassword compares the provided password with the stored hashed password
func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
