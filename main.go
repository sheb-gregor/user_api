package main

import (
	"log"
	"os"

	"user_api/cmd"
	"user_api/config"

	"github.com/urfave/cli"
)

func main() {
	appInfo := config.AppInfo()

	app := cli.NewApp()
	app.Usage = "A " + appInfo.Name + " service"
	app.Version = appInfo.Version + ", build " + appInfo.Build
	app.Flags = cmd.GetFlags()
	app.Commands = cmd.GetCommands()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
