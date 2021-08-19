package main

import (
	"github.com/locnguyenvu/mangden/internal/user"
	"github.com/locnguyenvu/mangden/pkg/config"
	"github.com/locnguyenvu/mangden/pkg/database"
	"github.com/locnguyenvu/mangden/pkg/logger"
	"github.com/locnguyenvu/mangden/pkg/view/template"
	"go.uber.org/dig"
)

func BuildApp() (*WebApp, error) {
	c := dig.New()

	constructors := []interface{}{
		config.New,
		database.New,
		database.NewGorm,
		logger.New,
		template.NewEngine,
		user.NewRepository,
		user.NewHttpHandler,
		NewRouter,
		NewWebApp,
	}

	for _, constructor := range constructors {
		if err := c.Provide(constructor); err != nil {
			return nil, err
		}
	}

	var webapp *WebApp
	err := c.Invoke(func(a *WebApp) {
		webapp = a
	})
	return webapp, err
}
