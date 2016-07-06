package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "google-apps-utils"
	app.Version = Version
	app.Author = "Hiroki Yagita"
	app.Usage = "Google Apps Utils with Directory API"
	app.Commands = Commands
	app.Run(os.Args)
}
