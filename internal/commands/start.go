package commands

import (
	"fmt"
	"log"

	"github.com/csyezheng/iChat/internal/context"
	"github.com/csyezheng/iChat/internal/server"
	"github.com/urfave/cli"
)

// Starts web server (user interface)
var StartCommand = cli.Command{
	Name:   "start",
	Usage:  "Starts web server",
	Flags:  startFlags,
	Action: startAction,
}

var startFlags = []cli.Flag{
	cli.IntFlag{
		Name:   "server-port, p",
		Usage:  "HTTP server port",
		Value:  80,
		EnvVar: "ICHAT_SERVER_PORT",
	},
	cli.StringFlag{
		Name:   "server-host, i",
		Usage:  "HTTP server host",
		Value:  "",
		EnvVar: "ICHAT_SERVER_HOST",
	},
	cli.StringFlag{
		Name:   "server-mode, m",
		Usage:  "debug, release or test",
		Value:  "",
		EnvVar: "ICHAT_SERVER_MODE",
	},
}

func startAction(ctx *cli.Context) error {
	conf := context.NewConfig(ctx)

	log.Print(conf)

	if conf.HttpServerPort() < 1 {
		log.Fatal(conf.HttpServerPort())
		log.Fatal("Server port must be a positive integer")
	}

	conf.MigrateDb()

	fmt.Printf("Starting web server at %s:%d...\n", ctx.String("server-host"), ctx.Int("server-port"))

	server.Start(conf)

	fmt.Println("Done.")

	return nil
}
