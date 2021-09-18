package user

import "time"

type User struct {
	resource user

	ID           int64
	Username     string
	PasswordHash string
	IsActive     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUser(resource user) *User {
	return &User{
		resource: resource,
	}
}

func (u *User) Resource() *user {
	return &u.resource
}
