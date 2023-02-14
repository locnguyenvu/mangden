package user

import "time"

type user struct {
	ID           int64
	Username     string `gorm:"uniqueIndex;size:256"`
	PasswordHash string
	IsActive     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func Orm() *user {
	return &user{}
}
