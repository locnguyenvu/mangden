package user

import (
	"time"

	modelhelper "github.com/locnguyenvu/mangden/pkg/database/model"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	resource user

	ID           int64
	Username     string
	PasswordHash string
	IsActive     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func newUser(orm user) *User {
	model := &User{
		resource: orm,
	}
	modelhelper.CopyFromOrm(model, &orm)
	return model
}

func (u *User) Resource() *user {
	return &u.resource
}

func (u *User) SetPassword(password string) *User {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.PasswordHash = string(hashPassword)
	return u
}

func (u *User) SetUsername(username string) *User {
	u.Username = username
	return u
}

func (u *User) Active() bool {
	return u.IsActive == 1
}
