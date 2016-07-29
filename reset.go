package main

import (
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

// reset ------------------------------------------------

var commandReset = cli.Command{
	Name:        "reset",
	Usage:       "Reset current oauth token",
	Description: "Reset current oauth token on Google Apps for Work",
	Action:      doReset,
}

func doReset(c *cli.Context) error {
	path, err := profilePath()
	if err != nil {
		log.Fatalf("Unable to get profile path: %v", err)
		return err
	}

	v010TokenPath := filepath.Join(path, url.QueryEscape("token.json"))
	os.Remove(v010TokenPath)

	dpath := filepath.Join(path, url.QueryEscape("directory-token.json"))
	os.Remove(dpath)

	rpath := filepath.Join(path, url.QueryEscape("reports-token.json"))
	os.Remove(rpath)

	return nil
}
