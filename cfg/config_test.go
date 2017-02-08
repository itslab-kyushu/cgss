//
// cfg/config_test.go
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
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func TestConfig(t *testing.T) {

	var err error
	config := Config{
		Groups: []Group{
			Group{
				Name: "group-1",
				Servers: []Server{
					Server{
						Address: "kvs1.group1.com",
						Port:    13009,
					},
					Server{
						Address: "kvs2.group1.com",
						Port:    13009,
					},
				},
			},
			Group{
				Name: "group-2",
				Servers: []Server{
					Server{
						Address: "kvs1.group2.com",
						Port:    13009,
					},
				},
			},
			Group{
				Name: "group-3",
				Servers: []Server{
					Server{
						Address: "kvs1.group3.com",
						Port:    13009,
					},
				},
			},
		},
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		t.Fatal(err.Error())
	}

	var res Config
	if err = yaml.Unmarshal(data, &res); err != nil {
		t.Fatal(err.Error())
	}

	if len(res.Groups) != 3 {
		t.Error("The number of groups is not correct:", res)
	}
	if len(res.Groups[0].Servers) != 2 {
		t.Error("The number of servers is not correct:", res)
	}

}
