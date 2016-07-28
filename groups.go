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
	Usage: "Shows current groups",
	Description: `Shows current groups on Google Apps for Work

			The record format:
			   [ID] EMAIL (NAME DESCRIPTION)
			`,
	Action: doGroupList,
}

func doGroupList(c *cli.Context) error {
	srv, err := NewService()
	if err != nil {
		return nil
	}

	r, err := srv.directory.Groups.List().Customer("my_customer").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve groups in domain. %v", err)
		return nil
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
	srv, err := NewService()
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

	group2, err := srv.directory.Groups.Insert(group).Do()
	if err != nil {
		log.Fatalf("Cannot create group in domain. %v", err)
	} else {
		log.Printf("Succeed to create group: %v", group2)
	}

	return nil
}

// group member list ----------------------------------------------

var commandGroupMemberList = cli.Command{
	Name:        "groupmemberlist",
	Usage:       "Shows a list of all members in a group",
	Description: "Shows a list of all members in a group on Google Apps for Work",
	Action:      doGroupMemberList,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "key, k", Value: "", Usage: "The group key"},
	},
}

func doGroupMemberList(c *cli.Context) error {
	srv, err := NewService()
	if err != nil {
		return nil
	}

	groupKey := c.String("key")

	r, err := srv.directory.Members.List(groupKey).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve group members in domain. %v", err)
		return nil
	}

	count := len(r.Members)
	if count == 0 {
		fmt.Print("No group members found.\n")
	} else {
		for _, m := range r.Members {
			fmt.Printf("%v, %v\n", m.Email, m.Role)
		}
	}

	return nil
}

// group create ----------------------------------------------

var commandGroupMemberCreate = cli.Command{
	Name:        "groupmembercreate",
	Usage:       "Add a member to the specified group",
	Description: "Add a member to the specified group on Google Apps for Work",
	Action:      doGroupMemberCreate,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "key, k", Value: "", Usage: "The group key"},
		cli.StringFlag{Name: "email, e", Value: "", Usage: "The member's email address"},
		cli.StringFlag{Name: "role, r", Value: "", Usage: "A group member's role can be: OWNER / MANAGER / MEMBER. The default is MEMBER, see: https://goo.gl/VIo3zH"},
	},
}

// Role is a group member role
type Role int

const (
	// DEFAULT equals MEMBER
	DEFAULT Role = iota
	// OWNER role can change send messages to the group, add or remove members, change member roles, change group's settings, and delete the group. An OWNER must be a member of the group.
	OWNER
	// MANAGER role is only available if the Google Groups for Business is enabled using the Admin console. A MANAGER role can do everything done by an OWNER role except make a member an OWNER or delete the group. A group can have multiple OWNER and MANAGER members.
	MANAGER
	// MEMBER role can subscribe to a group, view discussion archives, and view the group's membership list. For more information about member roles, see https://goo.gl/10gxHA.
	MEMBER
)

func (r Role) String() string {
	switch r {
	case OWNER:
		return "OWNER"
	case MANAGER:
		return "MANAGER"
	default:
		return "MEMBER"
	}
}

func roleFromString(role string) Role {
	switch {
	case role == "OWNER":
		return OWNER
	case role == "MANAGER":
		return MANAGER
	default:
		return MEMBER
	}
}

func doGroupMemberCreate(c *cli.Context) error {
	srv, err := NewService()
	if err != nil {
		return nil
	}

	groupKey := c.String("key")
	email := c.String("email")
	role := c.String("role")
	roleStr := roleFromString(role)

	member := &admin.Member{
		Email: email,
		Role:  roleStr.String(),
	}

	log.Printf("Key:%v, Email:%v, Role:%v[%v], Member:%v", groupKey, email, role, roleStr, member)

	member2, err := srv.directory.Members.Insert(groupKey, member).Do()
	if err != nil {
		log.Fatalf("Cannot add a member to the specified group in domain. %v", err)
		return nil
	}

	log.Printf("Succeed to add a member to the specified group: %v, member: %v", groupKey, member2.Email)

	return nil
}
