//
// client/command/simple/reconstruct.go
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
	if len(components) < 2 {
		return ""
	}
	return strings.Join(components[:len(components)-2], ".")

}
