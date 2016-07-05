package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gau"
	app.Version = Version
	app.Author = "Hiroki Yagita"
	app.Usage = "gau -- Google Apps Utils"
	app.Commands = Commands
	app.Run(os.Args)
}
