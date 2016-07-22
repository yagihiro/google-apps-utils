package main

import "github.com/urfave/cli"

// Commands is whole command objects
var Commands = []cli.Command{
	commandReset,
	commandList,
	commandCreate,
	commandGroupList,
	commandGroupCreate,
	commandGroupMemberList,
	commandGroupMemberCreate,
	// commandReports,
}
