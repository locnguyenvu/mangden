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
	"gorm.io/gorm/logger"
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

type EnvConfig struct {
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

func (c EnvConfig) DbConnectUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", c.DbUser, c.DbPassword, c.DbHost, c.DbPort, c.DbName)
}

func loadEnvConfig() *EnvConfig {
    cfg := &EnvConfig{
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
    return cfg
}


func NewGorm(applogger logrus.FieldLogger) (*gorm.DB, error) {
    cfg := loadEnvConfig()
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
