//
// client/command/remote/delete.go
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

package remote

import (
	"context"
	"fmt"
	"os"
	"runtime"

	pb "gopkg.in/cheggaaa/pb.v1"

	"golang.org/x/sync/errgroup"

	"github.com/itslab-kyushu/cgss/cfg"
	"github.com/itslab-kyushu/cgss/kvs"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

// CmdDelete prepares deleting a file and run cmdDelete.
func CmdDelete(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		return cli.ShowSubcommandHelp(c)
	}

	conf, err := cfg.ReadConfig(c.String("config"))
	if err != nil {
		return
	}
	return cmdDelete(conf, c.Args().First())

}

func cmdDelete(conf *cfg.Config, name string) (err error) {

	// Configure logging.
	bar := pb.New(conf.NServers())
	bar.Output = os.Stderr
	bar.Prefix("Server")
	bar.Start()
	defer bar.Finish()

	wg, ctx := errgroup.WithContext(context.Background())
	cpus := runtime.NumCPU()
	semaphore := make(chan struct{}, cpus)

	for _, group := range conf.Groups {

		for _, server := range group.Servers {

			semaphore <- struct{}{}
			func(server cfg.Server) {
				wg.Go(func() (err error) {
					defer func() { <-semaphore }()
					defer bar.Increment()

					conn, err := grpc.Dial(
						fmt.Sprintf("%s:%d", server.Address, server.Port),
						grpc.WithInsecure(),
						grpc.WithCompressor(grpc.NewGZIPCompressor()),
						grpc.WithDecompressor(grpc.NewGZIPDecompressor()),
					)
					if err != nil {
						return
					}
					defer conn.Close()

					client := kvs.NewKvsClient(conn)
					_, err = client.Delete(ctx, &kvs.Key{
						Name: name,
					})
					return

				})
			}(server)

		}

	}

	return wg.Wait()

}
