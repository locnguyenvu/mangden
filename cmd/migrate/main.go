package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"gorm.io/gorm"

	"github.com/locnguyenvu/mdn/internal/user"
	"github.com/locnguyenvu/mdn/pkg/database/mysql"
	"github.com/locnguyenvu/mdn/pkg/logging"
)


func main() {
    c := bootstrap()
    c.Invoke(func(db *gorm.DB, logger logrus.FieldLogger) {
        err := db.AutoMigrate(&user.User{})
        if err != nil {
            logger.Errorf("Failed to run auto migration", err)
        }
        logger.Info("Success!")
    })
}

func bootstrap() *dig.Container {
    c := dig.New()

    constructors := []interface{}{
        logging.NewFromEnv,
        mysql.NewGormFromEnv,
    }

    var err error
    for _, constructor := range constructors {
        if err = c.Provide(constructor); err != nil {
            fmt.Printf("Failed to bootstrap %T : %s", constructor, err.Error())
            panic(1)
        }
    }
    return c
}
