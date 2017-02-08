package remote

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"golang.org/x/sync/errgroup"

	"github.com/cheggaaa/pb"
	"github.com/itslab-kyushu/cgss/cfg"
	"github.com/itslab-kyushu/cgss/kvs"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

// CmdDelete prepares deleting a file and run cmdDelete.
func CmdDelete(c *cli.Context) (err error) {

	if c.NArg() != 2 {
		return cli.ShowSubcommandHelp(c)
	}

	conf, err := cfg.ReadConfig(c.Args().First())
	if err != nil {
		return
	}
	return cmdDelete(conf, c.Args().Get(1))

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
			func(server *cfg.Server) {
				wg.Go(func() (err error) {
					defer func() { <-semaphore }()
					defer bar.Increment()

					conn, err := grpc.Dial(fmt.Sprintf("%s:%d", server.Address, server.Port), grpc.WithInsecure())
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
			}(&server)

		}

	}

	return wg.Wait()

}
