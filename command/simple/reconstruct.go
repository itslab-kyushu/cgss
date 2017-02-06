package simple

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/itslab-kyushu/cgss/sss"
	"github.com/urfave/cli"
)

type reconstructOpt struct {
	ShareFiles []string
	OutputFile string
}

// CmdReconstruct executes reconstruct command.
func CmdReconstruct(c *cli.Context) error {

	if c.NArg() == 0 {
		return cli.ShowSubcommandHelp(c)
	}

	opt := &reconstructOpt{
		ShareFiles: append([]string{c.Args().First()}, c.Args().Tail()...),
		OutputFile: c.String("output"),
	}
	if opt.OutputFile == "" {
		opt.OutputFile = outputFile(opt.ShareFiles[0])
	}

	return cmdReconstruct(opt)

}

func cmdReconstruct(opt *reconstructOpt) error {

	shares := make([]sss.Share, len(opt.ShareFiles))
	for i, f := range opt.ShareFiles {

		data, err := ioutil.ReadFile(f)
		if err != nil {
			return err
		}

		if err = json.Unmarshal(data, &shares[i]); err != nil {
			return err
		}

	}

	secret, err := sss.Reconstruct(shares)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(opt.OutputFile, secret, 0644)

}

// outputFile returns a filename associated with the given share file name.
func outputFile(sharename string) string {

	components := strings.Split(sharename, ".")
	return strings.Join(components[:len(components)-2], ".")

}
