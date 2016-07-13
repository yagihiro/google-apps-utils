package main

import (
	"fmt"
	"log"

	"google.golang.org/api/admin/directory/v1"

	"github.com/urfave/cli"
)

// group list ----------------------------------------------

var commandGroupList = cli.Command{
	Name:  "grouplist",
	Usage: "Show current groups",
	Description: `Show current groups on Google Apps for Work

			The record format:
			   [ID] EMAIL (NAME DESCRIPTION)
			`,
	Action: doGroupList,
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
			fmt.Printf("[%v] %v (%v: %v)\n", g.Id, g.Email, g.Name, g.Description)
		}
	}

	return nil
}

// group create ----------------------------------------------

var commandGroupCreate = cli.Command{
	Name:        "groupcreate",
	Usage:       "Create a new group",
	Description: "Create a new group on Google Apps for Work",
	Action:      doGroupCreate,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "email, e", Value: "", Usage: "The group's email address"},
		cli.StringFlag{Name: "description, d", Value: "", Usage: "An extended description to help users determine the purpose of a group"},
		cli.StringFlag{Name: "name, n", Value: "", Usage: "The group's name"},
	},
}

func doGroupCreate(c *cli.Context) error {
	srv, err := GetService()
	if err != nil {
		return nil
	}

	email := c.String("email")
	description := c.String("description")
	name := c.String("name")

	log.Printf("Email:%v, Description:%v, Name:%v", email, description, name)

	group := &admin.Group{
		Email:       email,
		Description: description,
		Name:        name,
	}

	group2, err := srv.Groups.Insert(group).Do()
	if err != nil {
		log.Fatalf("Cannot create group in domain. %v", err)
	} else {
		log.Printf("Succeed to create group: %v", group2)
	}

	return nil
}
