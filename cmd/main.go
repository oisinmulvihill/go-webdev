package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/oisinmulvihill/go-webdev/internal/settings"
	"github.com/oisinmulvihill/go-webdev/internal/storage"
	"github.com/oisinmulvihill/go-webdev/internal/views"
)

func main() {
	configuration := settings.Recover(os.Args[1:])

	log.Printf("Setting up routing")
	routes := views.SetUp(configuration)

	log.Printf("Getting a database connection")
	conn, err := storage.Connection(configuration.DatabaseDSN)
	if err != nil {
		log.Printf("Unable to connect to the database, because: %s", err)
		os.Exit(1)
	}

	log.Printf("Setting up database")
	storage.Init(conn)

	log.Printf("Starting server on %s:%d", configuration.Interface, configuration.Port)
	log.Printf("Server version: %s", configuration.GitCommit)
	http.ListenAndServe(fmt.Sprintf("%s:%d", configuration.Interface, configuration.Port), routes)

	log.Printf("Web server exiting.")
	os.Exit(0)
}
