//
// client/command/simple/distribute.go
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

package simple

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

	threshold, err := strconv.Atoi(c.Args().Get(1))
	if err != nil {
		return
	}
	size, err := strconv.Atoi(c.Args().Get(2))
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
