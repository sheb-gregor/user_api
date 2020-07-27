package cmd

import (
	"user_api/config"

	"github.com/lancer-kit/uwe/v2"
	"github.com/urfave/cli"
)

func GetCommands() []cli.Command {
	return []cli.Command{
		serveCmd(),
		uwe.CliCheckCommand(config.AppInfo(), func(c *cli.Context) []uwe.WorkerName {
			return []uwe.WorkerName{config.WorkerAPIServer}
		}),
	}
}

func GetFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  config.FlagConfig + ", c",
			Value: "./config.yaml",
		},
	}
}
