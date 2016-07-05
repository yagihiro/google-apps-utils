package main

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	commandList,
}

var commandList = cli.Command{
	Name:        "list",
	Usage:       "Show current users",
	Description: "Show current users on Google Apps for Work",
	Action:      doList,
}

func doList(c *cli.Context) error {
	//	log.Printf("doList %v", c)

	srv, err := GetService()
	if err != nil {
		return nil
	}

	r, err := srv.Users.List().Customer("my_customer").MaxResults(10).OrderBy("email").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve users in domain. %v", err)
	}

	count := len(r.Users)
	if count == 0 {
		fmt.Print("No users found.\n")
	} else {
		for _, u := range r.Users {
			fmt.Printf("%s (%s)\n", u.PrimaryEmail, u.Name.FullName)
		}
	}

	return nil
}
