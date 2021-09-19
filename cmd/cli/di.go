package main

import (
	"github.com/locnguyenvu/mangden/internal/console"
	"github.com/locnguyenvu/mangden/internal/user"
	"github.com/locnguyenvu/mangden/pkg/config"
	"github.com/locnguyenvu/mangden/pkg/database"
	"github.com/locnguyenvu/mangden/pkg/logger"
	"go.uber.org/dig"
)

func bootstrap() (*dig.Container, error) {
	c := dig.New()

	constructors := []interface{}{
		config.New,
		logger.New,
		database.NewGorm,
		user.NewRepository,

		console.NewHandler,
		NewCommands,
	}

	for _, constructor := range constructors {
		if err := c.Provide(constructor); err != nil {
			return nil, err
		}
	}

	return c, nil
}
