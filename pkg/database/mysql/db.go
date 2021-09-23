package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/locnguyenvu/mangden/pkg/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGorm(cfg *config.Config, applogger logrus.FieldLogger) (*gorm.DB, error) {
	db, err := sql.Open("mysql", cfg.DbConnectUrl())
	if err != nil {
		applogger.Error(err)
		panic(err)
	}
	db.SetMaxIdleConns(cfg.DBMaxIdleConnections)
	db.SetMaxOpenConns(cfg.DBMaxOpenConnections)
	db.SetConnMaxLifetime(cfg.DBMaxConnLifetime)

	logConfig := logger.Config{
		Colorful:                  true,
		SlowThreshold:             time.Second,
		IgnoreRecordNotFoundError: true,
		LogLevel:                  logger.Silent,
	}
	gormLogger := logger.New(applogger, logConfig)
	return gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: gormLogger,
	})
}
