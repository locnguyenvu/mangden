package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

func main() {
	c, err := bootstrap()
	if err != nil {
		fmt.Println(err)
	}
	var l logrus.FieldLogger
	c.Invoke(func(logger logrus.FieldLogger) {
		l = logger
	})

	var app *cli.App
	c.Invoke(func(commands []*cli.Command) {
		app = &cli.App{
			Name:     "Mangden",
			Usage:    "make an explosive entrance",
			Commands: commands,
		}
	})
	appError := app.Run(os.Args)
	if appError != nil {
		l.Fatal(appError, "Failed to start daemon session")
	}
}
