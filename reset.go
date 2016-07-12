package main

import "github.com/urfave/cli"

// reset ------------------------------------------------

var commandReset = cli.Command{
	Name:        "reset",
	Usage:       "Reset current oauth token",
	Description: "Reset current oauth token on Google Apps for Work",
	Action:      doReset,
}

func doReset(c *cli.Context) error {
	ResetToken()
	return nil
}
