package settings

import (
	_ "embed"
	"flag"

	"github.com/oisinmulvihill/go-webdev/internal/core"
)

//go:generate sh -c "git rev-parse HEAD > commit.txt"
//go:embed commit.txt
var Commit string

func Recover(arguments []string) *core.Configuration {

	flags := flag.FlagSet{}
	bindInterface := flags.String("interface", "0.0.0.0", "The interface to bind to.")
	tcpPort := flags.Int("port", 8080, "The TCP port to listen on.")
	flags.Parse(arguments)

	config := core.Configuration{
		Interface: *bindInterface,
		Port:      *tcpPort,
		GitCommit: Commit,
	}

	return &config
}
