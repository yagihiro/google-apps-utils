package main

import (
	"fmt"
	"log"

	"google.golang.org/api/admin/directory/v1"

	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	commandList,
	commandCreate,
}

// list ------------------------------------------------

var commandList = cli.Command{
	Name:        "list",
	Usage:       "Show current users",
	Description: "Show current users on Google Apps for Work",
	Action:      doList,
}

func doList(c *cli.Context) error {
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

// create ------------------------------------------------

var commandCreate = cli.Command{
	Name:        "create",
	Usage:       "Create a new user",
	Description: "Create a new user on Google Apps for Work",
	Action:      doCreate,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "given, g", Value: "", Usage: "Given Name"},
		cli.StringFlag{Name: "family, f", Value: "", Usage: "Family Name"},
		cli.StringFlag{Name: "email, e", Value: "", Usage: "Primary Email"},
	},
}

func doCreate(c *cli.Context) error {
	srv, err := GetService()
	if err != nil {
		return nil
	}

	given := c.String("given")
	family := c.String("family")
	email := c.String("email")
	password := RandomString(8)

	log.Printf("Given:%v, Family:%v, Email:%v, Password:%v", given, family, email, password)

	name := &admin.UserName{
		GivenName:  given,
		FamilyName: family,
	}
	user := &admin.User{
		Name:                      name,
		Password:                  password,
		PrimaryEmail:              email,
		ChangePasswordAtNextLogin: true,
	}

	user2, err := srv.Users.Insert(user).Do()
	if err != nil {
		log.Fatalf("Cannot create user in domain. %v", err)
	} else {
		log.Fatalf("Succeed to create user: %v", user2)
	}

	return nil
}