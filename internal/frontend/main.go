package frontend

import (
	"fmt"
	"html/template"
	"io"
	"naivegateway/internal/config"
	"naivegateway/internal/logger"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var log = logger.Log
var cfg = config.GetConfig()

// NewCommand creates a new command for the cli
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "frontend",
		Short: "Starts the frontend service",
		Run: func(cmd *cobra.Command, args []string) {
			api := New()
			api.start()
		},
	}
	return cmd
}

// API is a struct that holds together the api dependencies
type API struct {
}

// New creates a new api object
func New() API {
	return API{}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Trace(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func (api *API) start() {
	connectionString := fmt.Sprintf(":%s", cfg.Frontend.Port)
	r := mux.NewRouter()

	index := "web/frontend/build/index.html"
	static := "web/frontend/build/static"
	templateIndex(index)
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(static))))
	r.PathPrefix("/").HandlerFunc(IndexHandler(index))
	r.Use(loggingMiddleware)
	log.Infof("API server listening on port %s", cfg.Frontend.Port)
	err := http.ListenAndServe(connectionString, r)
	if err != nil {
		log.Fatal(err)
	}
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}

func envToMap() (map[string]string, error) {
	envMap := make(map[string]string)
	var err error

	for _, v := range os.Environ() {
		split_v := strings.Split(v, "=")
		envMap[split_v[0]] = split_v[1]
	}

	return envMap, err
}

func templateIndex(indexPath string) {
	backup := fmt.Sprintf("%s.tpl", indexPath)
	if _, err := os.Stat(backup); os.IsNotExist(err) {
		copy(indexPath, backup)
	}
	tmpl := template.Must(template.ParseFiles(backup))
	f, _ := os.Create(indexPath)
	envMap, _ := envToMap()
	tmpl.Execute(f, envMap)
	f.Close()
}

func copy(source, dest string) {
	original, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	defer original.Close()

	new, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer new.Close()

	_, err = io.Copy(new, original)
	if err != nil {
		log.Fatal(err)
	}
}
