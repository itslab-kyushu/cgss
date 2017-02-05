package main

import (
	"fmt"
	"os"

	"github.com/itslab-kyushu/cgss/command"
	"github.com/urfave/cli"
)

// GlobalFlags defines a set of global flags.
var GlobalFlags = []cli.Flag{}

// Commands defines a set of commands.
var Commands = []cli.Command{
	{
		Name:        "distribute",
		Usage:       "Distribute a file",
		ArgsUsage:   "<file> <share size> <threshold>",
		Description: "distribute command makes a set of shares of a given file.",
		Action:      command.CmdDistribute,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "chunk",
				Usage: "Byte `size` of eash chunk.",
				Value: 256,
			},
		},
	},
	{
		Name:        "reconstruct",
		Usage:       "Reconstruct a file from a set of secrets",
		ArgsUsage:   "<file>...",
		Description: "reconstruct command reconstructs a file from a given set of shares.",
		Action:      command.CmdReconstruct,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "output",
				Usage: "Store the reconstructed secret to the `FILE`.",
			},
		},
	},
}

// CommandNotFound handles an error that the given command is not found.
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
