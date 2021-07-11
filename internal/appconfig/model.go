package appconfig

import (
	"time"

	"gorm.io/gorm"
)

type AppConfig struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	Name      string
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tabler interface {
	TableName() string
}

func (AppConfig) TableName() string {
	return "configs"
}

type Repository interface {
	DB() *gorm.DB
	Create(name string, value string)
	GetByName(name string) AppConfig
}

type repository struct {
	db *gorm.DB
}

func (r repository) Create(name string, value string) {
	config := &AppConfig{
		Name:  name,
		Value: value,
	}
	r.db.Save(config)
}

func (r repository) GetByName(name string) AppConfig {
	var result AppConfig
	r.db.First(&result, "name = ?", name)
	return result
}

func (r repository) DB() *gorm.DB {
	return r.db
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
