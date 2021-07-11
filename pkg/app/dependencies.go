package app

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Dependencies struct {
	Logger logrus.FieldLogger
	GormDB *gorm.DB
}
