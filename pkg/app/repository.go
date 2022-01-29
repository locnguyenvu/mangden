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

/**
 * Usage r.BulkLoadFromOrm(&[]*app.Model, []interface{})
 */
func (r Repository) BulkLoadFromOrm(models, orms interface{}) error {
	if reflect.TypeOf(models).Kind() != reflect.Ptr {
		return fmt.Errorf("invalid type")
	}
	if reflect.ValueOf(models).Elem().Kind() != reflect.Slice || reflect.TypeOf(orms).Kind() != reflect.Slice {
		return fmt.Errorf("invalid type")
	}
	// modelType = (elem of slice pointer).(elem of slice element)
	modelType := reflect.TypeOf(models).Elem().Elem()
	ormType := reflect.TypeOf(orms).Elem()

	loadfn := func(model, orm reflect.Value) {
		for i := 0; i < ormType.NumField(); i++ {
			fieldName := ormType.Field(i).Name
			model.Elem().FieldByName(fieldName).Set(orm.FieldByName(fieldName))
		}
		setResourceFn := model.Elem().FieldByName("Model").Addr().MethodByName("SetResource")
		setResourceFn.Call([]reflect.Value{orm})
	}

	container := reflect.ValueOf(models).Elem()
	for i := 0; i < reflect.ValueOf(orms).Len(); i++ {
		m := reflect.New(modelType.Elem())
		loadfn(m, reflect.ValueOf(orms).Index(i))
		container = reflect.Append(container, m)
	}
	reflect.ValueOf(models).Elem().Set(container)
	return nil
}

