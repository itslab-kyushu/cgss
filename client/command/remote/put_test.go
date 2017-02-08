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
