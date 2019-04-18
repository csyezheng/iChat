package commands

import (
	"fmt"

	"github.com/csyezheng/iChat/internal/context"
	"github.com/urfave/cli"
)

// Print current configuration
var ConfigCommand = cli.Command{
	Name:   "config",
	Usage:  "Displays global configuration values",
	Action: configAction,
}

func configAction(ctx *cli.Context) error {
	conf := context.NewConfig(ctx)

	fmt.Printf("NAME                  VALUE\n")
	fmt.Printf("debug                 %t\n", conf.Debug())
	fmt.Printf("config-file           %s\n", conf.ConfigFile())
	fmt.Printf("assets-path           %s\n", conf.AssetsPath())
	fmt.Printf("database-driver       %s\n", conf.DatabaseDriver())
	fmt.Printf("database-dsn          %s\n", conf.DatabaseDsn())

	return nil
}
