package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/locnguyenvu/mangden/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.DbConnectUrl())

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return db, nil
}

func NewGorm(dbConn *sql.DB) (*gorm.DB, error) {
	gormDb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: dbConn,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return gormDb, err
}
