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
