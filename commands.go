package main

import "github.com/urfave/cli"

var Commands = []cli.Command{
	commandReset,
	commandList,
	commandCreate,
	commandGroupList,
	commandGroupCreate,
}
