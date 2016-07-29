package main

import (
	"fmt"
	"log"
	"time"

	"github.com/urfave/cli"
)

// reports ------------------------------------------------

var commandReports = cli.Command{
	Name:        "reports",
	Usage:       "Show reports",
	Description: "Show activities, usage and audit reports on Google Apps for Work",
	Action:      doReports,
}

func doReports(c *cli.Context) error {
	srv, err := NewService()
	if err != nil {
		return nil
	}

	r, err := srv.reports.Activities.List("all", "login").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve logins to domain. %v", err)
	}

	if len(r.Items) == 0 {
		fmt.Println("No logins found.")
	} else {
		fmt.Println("Logins:")
		for _, a := range r.Items {
			t, err := time.Parse(time.RFC3339Nano, a.Id.Time)
			if err != nil {
				fmt.Println("Unable to parse login time.")
				// Set time to zero.
				t = time.Time{}
			}
			fmt.Printf("%s: %s %s\n", t.Format(time.RFC822), a.Actor.Email,
				a.Events[0].Name)
		}
	}

	return nil
}
