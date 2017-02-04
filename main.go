package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = Author
	app.Email = Email
	app.Usage = "Distribute and reconstruct a secret file."
	app.UsageText = "cgss command [command options] [arguments...]"

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.EnableBashCompletion = true
	app.Copyright = `cgss  Copyright (C) 2017  Junpei Kawamoto
This program comes with ABSOLUTELY NO WARRANTY.
This is free software, and you are welcome to redistribute it
under certain conditions.`

	app.Run(os.Args)
}
