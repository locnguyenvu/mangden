package model

import (
	"reflect"
	"regexp"
)

func CopyFromOrm(model, orm interface{}) {
	ormRv := reflect.ValueOf(orm).Elem()
	ormRvType := ormRv.Type()
	modelRv := reflect.ValueOf(model).Elem()

	for i := 0; i < ormRvType.NumField(); i++ {
		fieldName := ormRvType.Field(i).Name
		modelRv.FieldByName(fieldName).Set(ormRv.FieldByName(fieldName))
	}
}

func CopyFromModel(model, orm interface{}) {
	ormRv := reflect.ValueOf(orm).Elem()
	ormRvType := ormRv.Type()
	modelRv := reflect.ValueOf(model).Elem()

	for i := 0; i < ormRvType.NumField(); i++ {
		fieldName := ormRvType.Field(i).Name
		if fieldName == "ID" {
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
