package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/oisinmulvihill/go-webdev/internal/settings"
	"github.com/oisinmulvihill/go-webdev/internal/views"
)

func main() {
	configuration := settings.Recover(os.Args[1:])

	routes := views.SetUp(configuration)

	log.Printf("Starting server on %s:%d", configuration.Interface, configuration.Port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", configuration.Interface, configuration.Port), routes)

	log.Printf("Web server exiting.")
}
