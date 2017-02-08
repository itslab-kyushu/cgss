//
// client/command/remote/put_test.go
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

package remote

import (
	"testing"

	"github.com/itslab-kyushu/cgss/cfg"
)

func TestAllocation(t *testing.T) {

	config := cfg.Config{
		Groups: []cfg.Group{
			cfg.Group{
				Name: "group-1",
				Servers: []cfg.Server{
					cfg.Server{
						Address: "kvs1.group1.com",
						Port:    13009,
					},
					cfg.Server{
						Address: "kvs2.group1.com",
						Port:    13009,
					},
				},
			},
			cfg.Group{
				Name: "group-2",
				Servers: []cfg.Server{
					cfg.Server{
						Address: "kvs1.group2.com",
						Port:    13009,
					},
				},
			},
			cfg.Group{
				Name: "group-3",
				Servers: []cfg.Server{
					cfg.Server{
						Address: "kvs1.group3.com",
						Port:    13009,
					},
				},
			},
		},
	}

	res := allocation(&config)
	if len(res) != 3 {
		t.Error("Allocation returns a wrong allocation:", res)
	}
	if res[0] != 2 || res[1] != 1 || res[2] != 1 {
		t.Error("Allocation returns a wrong allocation:", res)
	}

}
