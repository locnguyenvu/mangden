package mysql

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/locnguyenvu/mangden/pkg/config"
)

func New(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.DbConnectUrl())

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return db, nil
}

func CreateOrm(cfg *config.Config) (*gorm.DB, error) {
	mysqlConn, _ := New(cfg)

	gormDb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: mysqlConn,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return gormDb, err
}
