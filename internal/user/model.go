package user

import (
	"time"

	"github.com/locnguyenvu/mangden/pkg/app"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	app.Model

	ID           int64
	Username     string
	PasswordHash string
	IsActive     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func newUser(orm user) *User {
	model := new(User)
	model.SetResource(&orm)
	return model
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
