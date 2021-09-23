package user

import (
	"fmt"

	"github.com/locnguyenvu/mangden/pkg/app"
	"gorm.io/gorm"
)

type Repository struct {
	app.Repository
}

func NewRepository(db *gorm.DB) *Repository {
	repository := new(Repository)
	repository.SetDB(db)
	return repository
}

func (r *Repository) Create(model *User) (*User, error) {
	row := user{
		Username:     model.Username,
		PasswordHash: model.PasswordHash,
	}

	result := r.DB().Create(&row)
	if result.Error != nil {
		return nil, result.Error
	}
	model.SetResource(row)
	return model, nil
}

func (r *Repository) Find(id int64) (*User, error) {
	orm := user{}
	result := r.DB().First(&orm, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Record not found id #%d", id)
	}

	model := newUser(orm)
	r.LoadFromOrm(model, &orm)
	return model, nil
}
