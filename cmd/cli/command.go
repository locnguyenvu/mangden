package main

import (
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
	Logger logrus.FieldLogger
}

func NewCommands(p CommandParam) []*cli.Command {
	router := &Router{}

	router.Register("test", func(ctx *cli.Context) error {
		p.Logger.Info("Hello, world")
		return nil
	})

	return router.Commands()
}
