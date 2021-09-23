package app

import (
	"reflect"
	"regexp"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository struct {
	db     *gorm.DB
	logger logrus.FieldLogger
}

func (r *Repository) SetDB(db *gorm.DB) {
	r.db = db
}

func (r *Repository) SetLogger(logger logrus.FieldLogger) {
	r.logger = logger
}

func (r *Repository) DB() *gorm.DB {
	return r.db
}

func (r *Repository) Logger() logrus.FieldLogger {
	return r.logger
}

func (r Repository) LoadFromOrm(model Modeler, orm interface{}) {
	ormRv := reflect.ValueOf(orm).Elem()
	ormRvType := ormRv.Type()
	modelRv := reflect.ValueOf(model).Elem()

	for i := 0; i < ormRvType.NumField(); i++ {
		fieldName := ormRvType.Field(i).Name
		modelRv.FieldByName(fieldName).Set(ormRv.FieldByName(fieldName))
	}
}

func (r Repository) SyncToOrm(model Modeler, orm interface{}) {
	ormRv := reflect.ValueOf(orm).Elem()
	ormRvType := ormRv.Type()
	modelRv := reflect.ValueOf(model).Elem()

	for i := 0; i < ormRvType.NumField(); i++ {
		fieldName := ormRvType.Field(i).Name
		if fieldName == "ID" || fieldName == "Id" {
			continue
		}
		gormTag, ok := ormRvType.Field(i).Tag.Lookup("gorm")
		if ok {
			matched, _ := regexp.MatchString("primaryKey", gormTag)
			if matched {
				continue
			}
		}
		ormRv.FieldByName(fieldName).Set(modelRv.FieldByName(fieldName))
	}
}

func (r *Repository) Update(model Modeler) error {
	resource := model.Resource()
	r.SyncToOrm(model, resource)
	result := r.DB().Save(resource)
	return result.Error
}
