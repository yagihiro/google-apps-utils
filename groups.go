package main

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
)

// group list ----------------------------------------------

var commandGroupList = cli.Command{
	Name:        "grouplist",
	Usage:       "Show current groups",
	Description: "Show current groups on Google Apps for Work",
	Action:      doGroupList,
}

func doGroupList(c *cli.Context) error {
	srv, err := GetService()
	if err != nil {
		return nil
	}

	r, err := srv.Groups.List().Customer("my_customer").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve groups in domain. %v", err)
	}

	count := len(r.Groups)
	if count == 0 {
		fmt.Print("No groups found.\n")
	} else {
		for _, g := range r.Groups {
			fmt.Printf("%v (%v: %v)\n", g.Email, g.Name, g.Description)
		}
	}

	return nil
}
