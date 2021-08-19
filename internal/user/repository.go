package user

import (
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(username, password string) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r repository) CreateUser(username, password string) (*User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	row := user{
		Username:     username,
		PasswordHash: string(hashPassword),
	}

	result := r.db.Create(&row)
	if result.Error != nil {
		return nil, result.Error
	}
	return NewUser(row), nil
}
