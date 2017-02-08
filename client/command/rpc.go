package command

import (
	"context"
	"fmt"
	"io"

	"github.com/itslab-kyushu/cgss/kvs"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

func CmdRpc(c *cli.Context) (err error) {

	conn, err := grpc.Dial("127.0.0.1:13009", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	client := kvs.NewKvsClient(conn)
	stream, err := client.List(context.Background(), &kvs.ListRequest{})
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
