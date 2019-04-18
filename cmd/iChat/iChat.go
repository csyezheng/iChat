package main

import (
	"os"

	"github.com/csyezheng/iChat/internal/commands"
	"github.com/urfave/cli"
)

var version = "development"

func main() {
	app := cli.NewApp()
	app.Name = "iChat"
	app.Usage = "instant message"
	app.Version = version
	app.Copyright = "(c) 2019 <github.com/csyezheng>"
	app.EnableBashCompletion = true
	app.Flags = commands.GlobalFlags

	app.Commands = []cli.Command{
		commands.ConfigCommand,
		commands.StartCommand,
		commands.MigrateCommand,
	}

	app.Run(os.Args)
}
