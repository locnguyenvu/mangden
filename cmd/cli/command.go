package main

import (
	"github.com/locnguyenvu/mangden/internal/console"
	"github.com/locnguyenvu/mangden/internal/user"
	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

type Router struct {
	commands []*cli.Command
}

func (ch *Router) Register(name string, handler func(ctx *cli.Context) error) {
	ch.commands = append(ch.commands, &cli.Command{
		Name:   name,
		Action: handler,
	})
}

func (ch Router) Commands() []*cli.Command {
	return ch.commands
}

type CommandParam struct {
	dig.In
	Logger         logrus.FieldLogger
	UserRepository *user.Repository
	Handler        *console.Handler
}

func NewCommands(p CommandParam) []*cli.Command {
	router := &Router{}

	router.Register("test", func(ctx *cli.Context) error {
		p.Logger.Info("Hello, world")
		return nil
	})

	router.Register("migrate", p.Handler.Migrate)
	router.Register("user:create", p.Handler.UserCreate)
	router.Register("user:info", p.Handler.UserInfo)
	router.Register("user:update", p.Handler.UserUpdate)

	return router.Commands()
}
