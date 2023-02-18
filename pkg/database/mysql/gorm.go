package mysql

import (
    "context"
	"database/sql"
    "fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
)

const (
    defaultLogLevel = "info"
	defaultLogFormat = "json"
	defaultDbPort = 3306
	defaultDBMaxIdleConnections = 10
	defaultDBMaxOpenConnections = 5
	defaultDBMaxConnLifeTime = 30 * time.Minute
)

type Config struct {
	DbHost string `config:"DB_HOST"`
	DbUser string `config:"DB_USER"`
	DbPassword string `config:"DB_PASSWORD"`
	DbPort int `config:"DB_PORT"`
	DbName string `config:"DB_NAME"`
	DBMaxIdleConnections int `config:"DB_MAX_IDLE_CONNECTIONS"`
	DBMaxOpenConnections int `config:"DB_MAX_OPEN_CONNECTIONS"`
	DBMaxConnLifetime time.Duration `config:"DB_MAX_CONN_LIFETIME"`
    LogFormat string `config:"LOG_FORMAT" validate:"oneof=text json"`
    LogLevel string `config:"LOG_LEVEL" validate:"oneof=debug info warn error fatal panic"`
}

func (c Config) DbConnectUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", c.DbUser, c.DbPassword, c.DbHost, c.DbPort, c.DbName)
}


func NewGormFromEnv(logger logrus.FieldLogger) (*gorm.DB, error) {
    cfg := &Config{
        LogLevel:             defaultLogLevel,
        LogFormat:            defaultLogFormat,
        DbPort:               defaultDbPort,
        DBMaxIdleConnections: defaultDBMaxIdleConnections,
        DBMaxOpenConnections: defaultDBMaxOpenConnections,
        DBMaxConnLifetime:    defaultDBMaxConnLifeTime,
    }

    ctx := context.Background()
    loader := confita.NewLoader(env.NewBackend())
    loader.Load(ctx, cfg)

	db, err := sql.Open("mysql", cfg.DbConnectUrl())
	if err != nil {
        panic(err)
	}
    if err = db.Ping(); err != nil {
        panic(err)
    }
    
	db.SetMaxIdleConns(cfg.DBMaxIdleConnections)
	db.SetMaxOpenConns(cfg.DBMaxOpenConnections)
	db.SetConnMaxLifetime(cfg.DBMaxConnLifetime)

	logConfig := gormlogger.Config{
		Colorful:                  true,
		SlowThreshold:             time.Second,
		IgnoreRecordNotFoundError: true,
		LogLevel:                  gormlogger.Silent,
	}
	gormLogger := gormlogger.New(logger, logConfig)
	return gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: gormLogger,
	})
}
