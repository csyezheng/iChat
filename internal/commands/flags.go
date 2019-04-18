package commands

import "github.com/urfave/cli"

// Global CLI flags
var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		Name:   "debug",
		Usage:  "run in debug mode",
		EnvVar: "ICHAT_DEBUG",
	},
	cli.StringFlag{
		Name:   "config-file, c",
		Usage:  "load configuration from `FILENAME`",
		Value:  "configs/iChat.yml",
		EnvVar: "ICHAT_CONFIG_FILE",
	},
}
