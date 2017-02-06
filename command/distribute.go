//
// command/distribute.go
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

package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/itslab-kyushu/cgss/cgss"
	"github.com/urfave/cli"
)

type distributeOpt struct {
	Filename       string
	ChunkSize      int
	Allocation     cgss.Allocation
	GroupThreshold int
	DataThreshold  int
}

// CmdDistribute executes distribute command.
func CmdDistribute(c *cli.Context) (err error) {

	// Arguments:
	// <file> <group threshold> <data threshold> <allocation>
	if c.NArg() < 4 {
		return cli.ShowSubcommandHelp(c)
	}

	gthreshold, err := strconv.Atoi(c.Args().Get(1))
	if err != nil {
		return
	}
	dthreshold, err := strconv.Atoi(c.Args().Get(2))
	if err != nil {
		return
	}

	var allocation cgss.Allocation
	rawAllocation := c.Args().Tail()[2:]
	if len(rawAllocation) == 1 {
		sp := strings.Split(rawAllocation[0], ",")
		allocation = make(cgss.Allocation, len(sp))
		for i, v := range sp {
			allocation[i], err = strconv.Atoi(v)
			if err != nil {
				return
			}
		}
	} else {
		allocation = make(cgss.Allocation, len(rawAllocation))
		for i, v := range rawAllocation {
			allocation[i], err = strconv.Atoi(v)
			if err != nil {
				return
			}
		}
	}

	return cmdDistribute(&distributeOpt{
		Filename:       c.Args().Get(0),
		ChunkSize:      c.Int("chunk"),
		Allocation:     allocation,
		GroupThreshold: gthreshold,
		DataThreshold:  dthreshold,
	})

}

func cmdDistribute(opt *distributeOpt) (err error) {

	fmt.Println(opt)

	secret, err := ioutil.ReadFile(opt.Filename)
	if err != nil {
		return
	}

	shares, err := cgss.Distribute(secret, opt.ChunkSize, opt.Allocation, opt.GroupThreshold, opt.DataThreshold)
	if err != nil {
		return
	}

	for _, s := range shares {

		data, err := json.Marshal(s)
		if err != nil {
			return err
		}

		g := s.GroupShare.Key.Text(16)
		d := s.DataShare.Key.Text(16)
		filename := fmt.Sprintf("%s.%s.%s.json", opt.Filename, g, d)
		if err = ioutil.WriteFile(filename, data, 0644); err != nil {
			return err
		}

	}
	return

}
