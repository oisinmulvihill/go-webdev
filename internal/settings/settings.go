package settings

import (
	"flag"
)

type configuration struct {
	Interface string
	Port      int
}

func Recover(arguments []string) *configuration {

	flags := flag.FlagSet{}
	bindInterface := flags.String("interface", "0.0.0.0", "The interface to bind to.")
	tcpPort := flags.Int("port", 8080, "The TCP port to listen on.")
	flags.Parse(arguments)

	config := configuration{
		Interface: *bindInterface,
		Port:      *tcpPort,
	}

	return &config
}
