package views

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oisinmulvihill/go-webdev/internal/core"
)

func siteRoot(config *core.Configuration, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Git Commit Hash: %s\n", config.GitCommit)
}

func bookPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("Book title '%s' and page '%s'", vars["title"], vars["page"])
}

func SetUp(config *core.Configuration) *mux.Router {

	log.Printf("Setting up URL handlers\n")

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		siteRoot(config, w, r)
	})
	r.HandleFunc("/books/{title}/page/{page}", bookPage)

	log.Printf("URL handlers have been set up.\n")

	return r
}
