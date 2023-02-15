package config

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
        "time"

	"gorm.io/gorm"
)

// Map to table `configs`
type config struct {
	ID        int64
	Name      string
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func AppConfigOrm() *config {
    return &config{}
}

type AppConfig struct {
	db *gorm.DB
}

func NewAppConfig(db *gorm.DB) *AppConfig {
	return &AppConfig{db}
}

func (r AppConfig) DB() *gorm.DB {
	return r.db
}

func (r AppConfig) NewOrUpdate(dbconfigname, dbconfigvalue string) error {
	var erc config
	result := r.db.Where("name = ?", dbconfigname).First(&erc)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	erc.Name = dbconfigname
	erc.Value = dbconfigvalue

	r.db.Save(&erc)
	return nil
}

func (r AppConfig) Load(v interface{}) error {
	var configNames []string
	var configValues []config
	rv := reflect.ValueOf(v).Elem()
	rvType := rv.Type()

	for i := 0; i < rv.NumField(); i++ {
		configName := rvType.Field(i).Tag.Get("dbconfigname")
		configNames = append(configNames, configName)
	}

	r.db.Find(&configValues, map[string]interface{}{"name": configNames})

	if rv.NumField() < len(configValues) {
		return errors.New("Invalid setting for config")
	}

	configDictionary := make(map[string]string)
	for _, dbcfg := range configValues {
		configDictionary[dbcfg.Name] = dbcfg.Value
	}

	for i := 0; i < rv.NumField(); i++ {
		dbconfigName := rvType.Field(i).Tag.Get("dbconfigname")
		dbconfigValue := configDictionary[dbconfigName]
		switch rvK := rvType.Field(i).Type.Kind(); rvK {
		case reflect.String:
			rv.Field(i).SetString(dbconfigValue)
		case reflect.Int:
		case reflect.Int8:
		case reflect.Int16:
		case reflect.Int32:
		case reflect.Int64:
			convInt, _ := strconv.ParseInt(dbconfigValue, 0, 0)
			rv.Field(i).SetInt(convInt)
		case reflect.Bool:
			convBool, _ := strconv.ParseBool(dbconfigValue)
			rv.Field(i).SetBool(convBool)
		}
	}
	return nil
}

func (r AppConfig) Update(dest interface{}, fields []string) error {
	var dbconfigName, dbconfigValue string
	rv := reflect.ValueOf(dest).Elem()
	rvType := rv.Type()
	for _, fna := range fields {
		stucField, ok := rvType.FieldByName(fna)
		if !ok {
			return errors.New(fmt.Sprintf("Field name %s does not exists", fna))
		}
		attrValueInterface := rv.FieldByName(fna).Interface()
		dbconfigName = stucField.Tag.Get("dbconfigname")

		switch stucField.Type.Kind() {
		case reflect.Int:
		case reflect.Int8:
		case reflect.Int16:
		case reflect.Int32:
		case reflect.Int64:
			attrValue, _ := attrValueInterface.(int64)
			dbconfigValue = fmt.Sprintf("%d", attrValue)
		case reflect.Bool:
			attrValue, _ := attrValueInterface.(bool)
			dbconfigValue = strconv.FormatBool(attrValue)
		case reflect.String:
			dbconfigValue = attrValueInterface.(string)
		}

		uerr := r.NewOrUpdate(dbconfigName, dbconfigValue)
		if uerr != nil {
			return uerr
		}
	}
	return nil
}

func (r AppConfig) Get(dbconfigname string) string {
	var row config
	if err := r.db.First(&row, "name = ?", dbconfigname).Error; err != nil {
		return ""
	}
	return row.Value
}
