package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/itslab-kyushu/cgss/sss"
	"github.com/urfave/cli"
)

type distributeOpt struct {
	Filename  string
	ChunkSize int
	Size      int
	Threshold int
}

// CmdDistribute executes distribute command.
func CmdDistribute(c *cli.Context) (err error) {

	if c.NArg() != 3 {
		return cli.ShowSubcommandHelp(c)
	}

	size, err := strconv.Atoi(c.Args().Get(1))
	if err != nil {
		return
	}
	threshold, err := strconv.Atoi(c.Args().Get(2))
	if err != nil {
		return
	}

	return cmdDistribute(&distributeOpt{
		Filename:  c.Args().Get(0),
		ChunkSize: c.Int("chunk"),
		Size:      size,
		Threshold: threshold,
	})
}

func cmdDistribute(opt *distributeOpt) (err error) {

	secret, err := ioutil.ReadFile(opt.Filename)
	if err != nil {
		return
	}

	shares, err := sss.Distribute(secret, opt.ChunkSize, opt.Size, opt.Threshold)
	if err != nil {
		return
	}

	for i, s := range shares {

		data, err := json.Marshal(s)
		if err != nil {
			return err
		}

		filename := fmt.Sprintf("%s.%d.json", opt.Filename, i)
		if err = ioutil.WriteFile(filename, data, 0644); err != nil {
			return err
		}

	}

	return nil
}
