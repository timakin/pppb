package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/timakin/pppb/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "inspect",
		Usage:  "",
		Action: command.CmdInspect,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
