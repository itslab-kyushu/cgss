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
