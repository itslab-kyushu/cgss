package cfg

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
