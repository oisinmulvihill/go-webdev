package settings

import (
	_ "embed"
	"flag"

	"github.com/oisinmulvihill/go-webdev/internal/core"
)

// var Commit = func() string {
// 	info, ok := debug.ReadBuildInfo()
// 	if ok {
// 		log.Printf("info=%v\n", info)
// 		for _, setting := range info.Settings {
// 			if setting.Key == "vcs.revision" {
// 				log.Printf("vcs.revision=%s\n", setting.Value)
// 				return setting.Value
// 			}
// 		}
// 	} else {
// 		log.Printf("build info fail=%v\n", ok)
// 	}
// 	return ""
// }()

//go:generate sh -c "printf %s $(git rev-parse HEAD) > commit.txt"
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
