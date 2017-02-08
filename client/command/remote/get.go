package remote

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"google.golang.org/grpc"

	"golang.org/x/sync/errgroup"

	"github.com/cheggaaa/pb"
	"github.com/itslab-kyushu/cgss/cfg"
	"github.com/itslab-kyushu/cgss/cgss"
	"github.com/itslab-kyushu/cgss/kvs"
	"github.com/urfave/cli"
)

type getOpt struct {
	Config      *cfg.Config
	Name        string
	OutputFile  string
	NConnection int
	Log         io.Writer
}

// CmdGet prepares get command and run cmdGet.
func CmdGet(c *cli.Context) (err error) {

	if c.NArg() != 2 {
		return cli.ShowSubcommandHelp(c)
	}

	conf, err := cfg.ReadConfig(c.Args().First())
	if err != nil {
		return
	}
	fmt.Println(conf)

	output := c.String("output")
	if output == "" {
		output = c.Args().Get(1)
	}

	var log io.Writer
	if c.GlobalBool("quiet") {
		log = ioutil.Discard
	} else {
		log = os.Stderr
	}

	return cmdGet(&getOpt{
		Config:      conf,
		Name:        c.Args().Get(1),
		OutputFile:  output,
		NConnection: c.Int("max-connection"),
		Log:         log,
	})

}

func cmdGet(opt *getOpt) (err error) {

	if opt.Config.NServers() == 0 {
		return fmt.Errorf("No server information is given.")
	}

	fmt.Fprintln(opt.Log, "Downloading shares")
	bar := pb.New(opt.Config.NServers())
	bar.Output = opt.Log
	bar.Prefix("Server")
	bar.Start()

	shares := make([]cgss.Share, opt.Config.NServers())
	wg, ctx := errgroup.WithContext(context.Background())
	semaphore := make(chan struct{}, opt.NConnection)
	var i int
	for _, group := range opt.Config.Groups {

		for _, server := range group.Servers {

			// Check the current context.
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			func(server *cfg.Server, i int) {

				semaphore <- struct{}{}
				wg.Go(func() (err error) {
					defer func() { <-semaphore }()
					defer bar.Increment()

					conn, err := grpc.Dial(fmt.Sprintf("%s:%d", server.Address, server.Port), grpc.WithInsecure())
					if err != nil {
						return
					}
					defer conn.Close()

					client := kvs.NewKvsClient(conn)
					value, err := client.Get(ctx, &kvs.Key{
						Name: opt.Name,
					})
					if err != nil {
						return
					}

					share, err := FromValue(value)
					if err != nil {
						return
					}
					shares[i] = *share
					return

				})

			}(&server, i)
			i++

		}

	}

	err = wg.Wait()
	bar.Finish()
	if err != nil {
		return
	}

	fmt.Fprintln(opt.Log, "Reconstructing the secret")
	secret, err := cgss.Reconstruct(context.Background(), shares, opt.Log)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(opt.OutputFile, secret, 0644)

}
