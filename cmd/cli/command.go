package main

import (
	"github.com/locnguyenvu/mangden/internal/console"
	"github.com/locnguyenvu/mangden/pkg/config"
	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

type Router struct {
	commands []*cli.Command
}

func (ch *Router) Register(name string, handler func(ctx *cli.Context) error, flags ...cli.Flag) {
	cmd := &cli.Command{
		Name:   name,
		Action: handler,
	}
	if len(flags) > 0 {
		cmd.Flags = make([]cli.Flag, 0)
		for _, fl := range flags {
			cmd.Flags = append(cmd.Flags, fl)
		}
	}
	ch.commands = append(ch.commands, cmd)
}

func (ch Router) Commands() []*cli.Command {
	return ch.commands
}

type CommandParam struct {
	dig.In
	Logger  logrus.FieldLogger
	Handler *console.Handler
    AppConfig *config.AppConfig
}

func NewCommands(p CommandParam) []*cli.Command {
	router := &Router{}

	router.Register("test", func(ctx *cli.Context) error {
		p.Logger.Info("Hello, world")
		return nil
	})

    router.Register("test-appconfig", func(ctx *cli.Context) error {
        aconfig := struct {
            Hostname    string `dbconfigname:"hostname"`
        }{}
        p.AppConfig.Load(&aconfig)
        p.Logger.Infof("Config value: %s | %T", aconfig.Hostname, aconfig.Hostname)
        return nil
    })

	router.Register("migrate", p.Handler.Migrate)

	router.Register("user:create", p.Handler.UserCreate)
	router.Register("user:info", p.Handler.UserInfo)
	router.Register("user:update", p.Handler.UserUpdate)


    // Config
    router.Register("config:set",
        func(ctx *cli.Context) error {
            name := ctx.String("name")
            value := ctx.String("value")
            return p.AppConfig.NewOrUpdate(name, value)
        }, 
        &cli.StringFlag{Name: "name", Required: true},
        &cli.StringFlag{Name: "value", Required: true},
    )
	router.Register("config:get",
        func (ctx *cli.Context) error {
            name := ctx.String("name")
            p.Logger.Info(p.AppConfig.Get(name))
            return nil
        },
		&cli.StringFlag{Name: "name", Required: true},
	)

	return router.Commands()
}
