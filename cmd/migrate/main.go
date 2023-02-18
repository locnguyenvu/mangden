package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"gorm.io/gorm"

	"mck.co/fuel/internal/user"
	"mck.co/fuel/pkg/database/mysql"
	"mck.co/fuel/pkg/logging"
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
