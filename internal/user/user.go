package user

import (
    "time"

    "golang.org/x/crypto/bcrypt"
)

type User struct {

	ID int64
	Username string `gorm:"uniqueIndex;size:256"`
    FirstName string `gorm:"not null;"`
    LastName string `gorm:"not null;"`
    Yob int `gorm:"not null;"`
	PasswordHash string
    IsActive int `gorm:"default:1;"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (u *User) SetPassword(password string) *User {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.PasswordHash = string(hashPassword)
	return u
}
