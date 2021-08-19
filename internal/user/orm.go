package user

import "time"

type user struct {
	ID           int64
	Username     string
	PasswordHash string
	IsActive     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
