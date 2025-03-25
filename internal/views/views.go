package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/oisinmulvihill/go-webdev/internal/core"
	"github.com/oisinmulvihill/go-webdev/internal/storage"
)

func siteRoot(gitCommit string, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The server us running OK.\nGit commit hash: %s\n", gitCommit)
}

type SystemUsers struct {
	Users []core.User
}

func showUsers(config *core.Configuration, template *template.Template, w http.ResponseWriter, r *http.Request) {

	systemUsers := SystemUsers{}

	log.Printf("Getting a database connection")
	conn, err := storage.Connection(config.DatabaseDSN)
	if err != nil {
		log.Printf("Unable to connect to the database, because: %s", err)
	}

	systemUsers.Users, err = storage.GetUsers(conn)
	if err != nil {
		log.Printf("Unable to connect to the database, because: %s", err)
	}
	log.Printf("Got users: %v", systemUsers.Users)

	template.Execute(w, systemUsers)
}

func SetUp(config *core.Configuration) *mux.Router {

	log.Printf("Setting up URL handlers\n")

	r := mux.NewRouter()

	layoutTmpl := template.Must(template.ParseFiles(filepath.Join(config.TemplateDir, "layout.html")))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		siteRoot(config.GitCommit, w, r)
	})

	r.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		showUsers(config, layoutTmpl, w, r)
	})

	log.Printf("URL handlers have been set up.\n")

	return r
}
