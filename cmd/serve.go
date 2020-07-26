package cmd

import (
	"user_api/api"
	"user_api/cmd/modules"
	"user_api/config"

	"github.com/lancer-kit/armory/log"
	"github.com/lancer-kit/uwe/v2"
	"github.com/urfave/cli"
)

func serveCmd() cli.Command {
	var serveCommand = cli.Command{
		Name:   "serve",
		Usage:  "starts " + config.ServiceName + " workers",
		Action: serveAction,
	}
	return serveCommand
}

func serveAction(c *cli.Context) error {
	cfg, err := modules.Init(c)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	logger := log.Get().WithField("app", config.ServiceName)


	chief := uwe.NewChief()
	chief.UseDefaultRecover()
	chief.EnableServiceSocket(config.AppInfo())
	chief.SetEventHandler(uwe.LogrusEventHandler(logger))

	logger = logger.WithField("app_layer", "workers")
	chief.AddWorker(config.WorkerAPIServer,
		api.GetServer(cfg, logger.WithField("worker", config.WorkerAPIServer)))


	return nil
}
