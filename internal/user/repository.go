package user

import (
	"fmt"

	modelhelper "github.com/locnguyenvu/mangden/pkg/database/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) DB() *gorm.DB {
	return r.db
}

func (r *Repository) Create(model *User) (*User, error) {
	row := user{
		Username:     model.Username,
		PasswordHash: model.PasswordHash,
	}

	result := r.db.Create(&row)
	if result.Error != nil {
		return nil, result.Error
	}
	model.resource = row
	return model, nil
}

func (r *Repository) Update(model *User) error {
	orm := model.Resource()
	modelhelper.CopyFromModel(model, orm)
	result := r.db.Save(orm)
	return result.Error
}

func (r *Repository) Find(id int64) (*User, error) {
	orm := user{}
	result := r.db.First(&orm, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Record not found id #%d", id)
	}

	return newUser(orm), nil
}
