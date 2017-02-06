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
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/itslab-kyushu/cgss/cgss"
	"github.com/urfave/cli"
)

type distributeOpt struct {
	cgss.DistributeOpt
	Filename string
	Dir      string
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
		Filename: c.Args().Get(0),
		Dir:      c.String("dir"),
		DistributeOpt: cgss.DistributeOpt{
			ChunkSize:      c.Int("chunk"),
			Allocation:     allocation,
			GroupThreshold: gthreshold,
			DataThreshold:  dthreshold,
		},
	})

}

func cmdDistribute(opt *distributeOpt) (err error) {

	secret, err := ioutil.ReadFile(opt.Filename)
	if err != nil {
		return
	}

	ctx := context.Background()
	shares, err := cgss.Distribute(ctx, secret, &opt.DistributeOpt)
	if err != nil {
		return
	}

	base := filepath.FromSlash(filepath.Join(filepath.ToSlash(opt.Dir), filepath.Base(filepath.ToSlash(opt.Filename))))

	wg, ctx := errgroup.WithContext(ctx)
	cpus := runtime.NumCPU()
	semaphore := make(chan struct{}, cpus)
	for _, s := range shares {

		// Check the current context.
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		func(s cgss.Share) {
			semaphore <- struct{}{}
			wg.Go(func() (err error) {
				defer func() { <-semaphore }()

				select {
				case <-ctx.Done():
					return
				default:
				}

				data, err := json.Marshal(s)
				if err != nil {
					return
				}

				g := s.GroupKey().Text(16)
				d := s.DataKey().Text(16)
				return ioutil.WriteFile(fmt.Sprintf("%s.%s.%s.json", base, g, d), data, 0644)

			})
		}(s)

	}
	return wg.Wait()

}
