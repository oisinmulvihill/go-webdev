package settings

import (
	_ "embed"
	"flag"
	"os"

	"github.com/oisinmulvihill/go-webdev/internal/core"
)

//go:generate sh -c "git rev-parse HEAD > commit.txt"
//go:embed commit.txt
var Commit string

func Recover(arguments []string) *core.Configuration {

	envDatabaseDSN := os.Getenv("DATABASE_DSN")
	if envDatabaseDSN == "" {
		envDatabaseDSN = "postgres://service:service@db:7432/webdev"
	}

	envTemplateDir := os.Getenv("TEMPLATE_DIR")
	if envTemplateDir == "" {
		envTemplateDir = "internal/templates"
	}

	flags := flag.FlagSet{}
	databaseDSN := flags.String("database-dsn", envDatabaseDSN, "The database to connect to.")
	templateDir := flags.String("templateDir", envTemplateDir, "The templates location directory.")
	bindInterface := flags.String("interface", "0.0.0.0", "The interface to bind to.")
	tcpPort := flags.Int("port", 8080, "The TCP port to listen on.")
	flags.Parse(arguments)

	config := core.Configuration{
		DatabaseDSN: *databaseDSN,
		Interface:   *bindInterface,
		Port:        *tcpPort,
		GitCommit:   Commit,
		TemplateDir: *templateDir,
	}

	return &config
}
