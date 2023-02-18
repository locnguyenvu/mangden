package user

import (
    "gorm.io/gorm"
)

type Repository struct {
    DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
    return &Repository{
        DB: db,
    }
}

func (r Repository) Create(username string, password string, firstname string, lastname string, yob int) (*User, error) {
    user := &User{
        Username: username,
        FirstName: firstname,
        LastName: lastname,
        Yob: yob,
    }
    user.SetPassword(password)
    var err error
    if cm := r.DB.Save(user); cm.Error != nil {
       err = cm.Error 
    }
    return user, err
}

func (r Repository) ListLatest() ([]User, error) {
    var rows []User
    result := r.DB.Order("created_at desc").Limit(10).Find(&rows)
    if result.Error != nil {
        return rows, result.Error
    }
    return rows, nil
}

func (r Repository) Get(id int64) *User {
    var row User
    result := r.DB.First(&row, id)
    if result.Error != nil {
        return nil
    }
    return &row
}

func (r Repository) Delete(id int64) error {
    result := r.DB.Delete(&User{}, id)
    return result.Error
}

func (r Repository) Save(model *User) error {
    result := r.DB.Save(model)
    return result.Error
}
