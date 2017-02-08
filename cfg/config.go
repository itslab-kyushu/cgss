//
// cfg/config.go
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

package cfg

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Config defines a set of group servers to distribute shares.
type Config struct {
	Groups []Group
}

// Server defines server information.
type Server struct {
	Address string
	Port    int
}

// Group defines group information.
type Group struct {
	Name    string
	Servers []Server
}

// ReadConfig reads a YAML formatted config file.
func ReadConfig(filename string) (conf *Config, err error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	conf = new(Config)
	err = yaml.Unmarshal(data, conf)
	return

}

// NGroups returns the number of groups.
func (c *Config) NGroups() int {
	return len(c.Groups)
}

// NServers returns the number of all servers.
func (c *Config) NServers() (n int) {
	for _, g := range c.Groups {
		n += len(g.Servers)
	}
	return
}
