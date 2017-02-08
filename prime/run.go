package main

import (
	"crypto/rand"
	"fmt"
	"strconv"

	"github.com/urfave/cli"
)

// CmdRun generates a random prime which has a given bit-length.
func CmdRun(c *cli.Context) (err error) {

	if c.NArg() == 0 {
		return fmt.Errorf("Bit length is not given")
	}

	length, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return
	}

	prime, err := rand.Prime(rand.Reader, length)
	if err != nil {
		return err
	}

	fmt.Println(prime.Text(16))
	return

}
