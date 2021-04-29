package api

import (
	"fmt"
	"naivegateway/internal/config"
	"naivegateway/internal/database"
	"naivegateway/internal/logger"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var log = logger.Log
var cfg = config.GetConfig()

// NewCommand creates a new command for the cli
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api",
		Short: "Starts the api service",
		Run: func(cmd *cobra.Command, args []string) {
			api := New()
			api.start()
		},
	}
	return cmd
}

// API is a struct that holds together the api dependencies
type API struct {
	db *pg.DB
}

// New creates a new api object
func New() API {
	return API{
		db: database.NewConnection(),
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Trace(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func (api *API) start() {
	connectionString := fmt.Sprintf(":%s", cfg.API.Port)
	r := mux.NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"})
	originsOk := handlers.AllowedOrigins(cfg.API.CorsOrigins)
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// Generic routes
	r.HandleFunc("/health", api.health)
	// Account routes
	r.HandleFunc("/v1/accounts/create", api.createAccount).Methods("POST")
	r.HandleFunc("/v1/accounts/deposit", api.depositToAccount).Methods("POST")
	r.HandleFunc("/v1/accounts/detail", api.accountDetails).Methods("POST")
	r.HandleFunc("/v1/accounts/statement", api.accountStatement).Methods("POST")
	// Transactions
	r.HandleFunc("/v1/transactions", api.getTransactions).Methods("GET")
	r.HandleFunc("/v1/transactions/create", api.createTransaction).Methods("POST")
	r.HandleFunc("/v1/transactions/execute", api.executeTransaction).Methods("POST")

	r.Use(loggingMiddleware)
	log.Infof("API server listening on port %s", cfg.API.Port)
	http.ListenAndServe(connectionString, handlers.CORS(originsOk, headersOk, methodsOk)(r))
}
