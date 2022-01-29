package app

import (
	"fmt"
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

func (r *Repository) Save(model Modeler) error {
	err := r.SyncToOrm(model)
	if err != nil {
		return err
	}
	result := r.db.Save(model.Resource())
	if result.Error != nil {
		r.LoadFromOrm(model, nil)
	}
	return result.Error
}

func (r *Repository) Update(model Modeler) error {
	resource := model.Resource()
	err := r.SyncToOrm(model)
	if err != nil {
		return err
	}
	result := r.db.Save(resource)
	return result.Error
}

func (r Repository) LoadFromOrm(model Modeler, orm interface{}) {
	if orm == nil {
		orm = model.Resource()
	}
	ormRv := reflect.ValueOf(orm).Elem()
	ormRvType := ormRv.Type()
	modelRv := reflect.ValueOf(model).Elem()

	for i := 0; i < ormRvType.NumField(); i++ {
		fieldName := ormRvType.Field(i).Name
		modelRv.FieldByName(fieldName).Set(ormRv.FieldByName(fieldName))
	}
	if model.Resource() == nil {
		model.SetResource(orm)
	}
}

func (r Repository) SyncToOrm(model Modeler) error {
	orm := model.Resource()
	if orm == nil {
		return fmt.Errorf("empty model - no resource found")
	}
	ormRv := reflect.ValueOf(orm).Elem()
	ormRvType := ormRv.Type()
	modelRv := reflect.ValueOf(model).Elem()

	primaryKeyTagSearch, _ := regexp.Compile("primaryKey")
	for i := 0; i < ormRvType.NumField(); i++ {
		fieldName := ormRvType.Field(i).Name
		if fieldName == "ID" || fieldName == "Id" {
			continue
		}
		gormTag, ok := ormRvType.Field(i).Tag.Lookup("gorm")
		if ok {
			if primaryKeyTagSearch.Match([]byte(gormTag)) {
				continue
			}
		}
		ormRv.FieldByName(fieldName).Set(modelRv.FieldByName(fieldName))
	}
	return nil
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
