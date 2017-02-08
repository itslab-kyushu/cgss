package remote

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"

	"google.golang.org/grpc"

	"github.com/itslab-kyushu/cgss/cfg"
	"github.com/itslab-kyushu/cgss/kvs"
	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

// CmdList prepares list command and run cmdList.
func CmdList(c *cli.Context) (err error) {

	if c.NArg() != 1 {
		return cli.ShowSubcommandHelp(c)
	}

	data, err := ioutil.ReadFile(c.Args().First())
	if err != nil {
		return
	}
	var conf cfg.Config
	if err = yaml.Unmarshal(data, &conf); err != nil {
		return
	}
	return cmdList(&conf)

}

func cmdList(conf *cfg.Config) (err error) {

	if len(conf.Groups) == 0 {
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
