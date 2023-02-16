package user

import (
    "time"

    "golang.org/x/crypto/bcrypt"

)

type User struct {
	ID           int64
	Username     string `gorm:"uniqueIndex;size:256"`
	PasswordHash string
	IsActive     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) SetPassword(password string) *User {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.PasswordHash = string(hashPassword)
	return u
}
