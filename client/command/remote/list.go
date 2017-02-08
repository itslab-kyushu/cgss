//
// client/command/remote/list.go
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
	"io"
	"math/rand"

	"google.golang.org/grpc"

	"github.com/itslab-kyushu/cgss/cfg"
	"github.com/itslab-kyushu/cgss/kvs"
	"github.com/urfave/cli"
)

// CmdList prepares list command and run cmdList.
func CmdList(c *cli.Context) (err error) {

	if c.NArg() != 0 {
		return cli.ShowSubcommandHelp(c)
	}

	conf, err := cfg.ReadConfig(c.String("config"))
	if err != nil {
		return
	}
	return cmdList(conf)

}

func cmdList(conf *cfg.Config) (err error) {

	if conf.NGroups() == 0 {
		return fmt.Errorf("No groups given: %v", conf)
	}

	group := conf.Groups[rand.Intn(len(conf.Groups))]
	if len(group.Servers) == 0 {
		return fmt.Errorf("Group %v doesn't have servers", group.Name)
	}

	server := group.Servers[rand.Intn(len(group.Servers))]
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", server.Address, server.Port), grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	ctx := context.Background()
	client := kvs.NewKvsClient(conn)
	stream, err := client.List(ctx, &kvs.ListRequest{})
	if err != nil {
		return
	}

	for {
		item, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		fmt.Println(item.Name)
	}

	return

}
