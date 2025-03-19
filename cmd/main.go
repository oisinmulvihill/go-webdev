package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	flags := flag.FlagSet{}
	httpPort := flags.Int("port", 8080, "The TCP port to listen on.")
	httpInterface := flags.String("interface", "0.0.0.0", "The TCP interface to bind to.")
	flags.Parse(os.Args[1:])

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	log.Printf("Starting server on %s:%d", *httpInterface, *httpPort)
	http.ListenAndServe(fmt.Sprintf("%s:%d", *httpInterface, *httpPort), nil)

	log.Printf("Web server exiting.")
}
