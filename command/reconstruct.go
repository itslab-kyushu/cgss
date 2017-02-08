//
// command/reconstruct.go
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
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/itslab-kyushu/cgss/cgss"
	"github.com/ulikunitz/xz"
	"github.com/urfave/cli"
)

type reconstructOpt struct {
	ShareFiles []string
	OutputFile string
	Log        io.Writer
}

// CmdReconstruct executes reconstruct command.
func CmdReconstruct(c *cli.Context) error {

	if c.NArg() == 0 {
		return cli.ShowSubcommandHelp(c)
	}

	var log io.Writer
	if c.GlobalBool("quiet") {
		log = ioutil.Discard
	} else {
		log = os.Stderr
	}

	opt := &reconstructOpt{
		ShareFiles: append([]string{c.Args().First()}, c.Args().Tail()...),
		OutputFile: c.String("output"),
		Log:        log,
	}
	if opt.OutputFile == "" {
		opt.OutputFile = outputFile(opt.ShareFiles[0])
	}

	return cmdReconstruct(opt)

}

func cmdReconstruct(opt *reconstructOpt) (err error) {

	fmt.Fprintln(opt.Log, "Reading share files")
	shares := make([]cgss.Share, len(opt.ShareFiles))
	for i, f := range opt.ShareFiles {

		var data []byte
		if strings.HasSuffix(f, ".xz") {
			fp, err := os.Open(f)
			if err != nil {
				return err
			}
			defer fp.Close()

			r, err := xz.NewReader(fp)
			if err != nil {
				return err
			}
			data, err = ioutil.ReadAll(r)
			if err != nil {
				return err
			}

		} else {
			data, err = ioutil.ReadFile(f)
			if err != nil {
				return err
			}
		}
		if err = json.Unmarshal(data, &shares[i]); err != nil {
			return err
		}
	}

	fmt.Fprintln(opt.Log, "Reconstructing the secret")
	secret, err := cgss.Reconstruct(context.Background(), shares, opt.Log)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(opt.OutputFile, secret, 0644)

}

// outputFile returns a filename associated with the given share file name.
func outputFile(sharename string) string {

	components := strings.Split(sharename, ".")
	if len(components) < 3 {
		return ""
	}
	return strings.Join(components[:len(components)-3], ".")

}
