//
// client/commands.go
//
// Copyright (c) 2017 Junpei Kawamoto
//
// This file is part of cgss.
//
// cgss is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// cgss is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with cgss.  If not, see <http://www.gnu.org/licenses/>.
//

package main

import (
	"fmt"
	"os"

	"github.com/itslab-kyushu/cgss/client/command/local"
	"github.com/itslab-kyushu/cgss/client/command/remote"
	"github.com/itslab-kyushu/cgss/client/command/simple"
	"github.com/urfave/cli"
)

// GlobalFlags defines a set of global flags.
var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "quiet",
		Usage: "Omit logging outputs.",
	},
}

// Commands defines a set of commands.
var Commands = cli.Commands{
	{
		Name:  "remote",
		Usage: "Access remote CGSS servers",
		Subcommands: cli.Commands{
			{
				Name:  "get",
				Usage: "Download shares and reconstruct a file",
			},
			{
				Name:  "put",
				Usage: "Distribute and store shares",
			},
			{
				Name:        "delete",
				Usage:       "Delete a file from all servers",
				Description: "Delete a file from all servers.",
				ArgsUsage:   "<config file> <file name>",
				Action:      remote.CmdDelete,
			},
			{
				Name:        "list",
				Usage:       "Get a list of files stored in servers",
				Description: "Receive a list of files stored in a random server.",
				ArgsUsage:   "<config file>",
				Action:      remote.CmdList,
			},
		},
	},
	{
		Name:  "local",
		Usage: "Run local file based CGSS scheme",
		Subcommands: cli.Commands{
			{
				Name:        "distribute",
				Usage:       "Distribute a file by Cross-Group Secret Sharing scheme",
				ArgsUsage:   "<file> <group threshold> <data threshold> <allocation>",
				Description: "distribute command makes a set of shares of a given file.",
				Action:      local.CmdDistribute,
				Flags: []cli.Flag{
					cli.IntFlag{
						Name:  "chunk",
						Usage: "Byte `size` of eash chunk.",
						Value: 256,
					},
					cli.StringFlag{
						Name:  "dir",
						Usage: "Store shares into the given directory.",
					},
					cli.BoolFlag{
						Name:  "no-compress",
						Usage: "Store shares without compression.",
					},
				},
			},
			{
				Name:        "reconstruct",
				Usage:       "Reconstruct a file from secrets by Cross-Group Secret Sharing scheme",
				ArgsUsage:   "<file>...",
				Description: "reconstruct command reconstructs a file from a given set of shares.",
				Action:      local.CmdReconstruct,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "output",
						Usage: "Store the reconstructed secret to the `FILE`.",
					},
				},
			},
		},
	},
	{
		Name:  "simple",
		Usage: "Run Shamir's Secret Sharing scheme",
		// Desctiprion: "",
		Subcommands: cli.Commands{
			{
				Name:        "distribute",
				Usage:       "Distribute a file",
				ArgsUsage:   "<file> <threshold> <share size>",
				Description: "distribute command makes a set of shares of a given file.",
				Action:      simple.CmdDistribute,
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
				Action:      simple.CmdReconstruct,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "output",
						Usage: "Store the reconstructed secret to the `FILE`.",
					},
				},
			},
		},
	},
}

// CommandNotFound handles an error that the given command is not found.
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
