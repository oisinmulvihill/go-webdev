package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/oisinmulvihill/go-webdev/internal/settings"
)

func main() {
	configuration := settings.Recover(os.Args[1:])

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	log.Printf("Starting server on %s:%d", configuration.Interface, configuration.Port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", configuration.Interface, configuration.Port), nil)

	log.Printf("Web server exiting.")
}
