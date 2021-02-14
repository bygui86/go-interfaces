package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/bygui86/go-testing/rest-examples/http-server/database"
)

type Server struct {
	config     *Config
	router     *mux.Router
	httpServer *http.Server
	db         database.InMemoryDb
	running    bool
}

type Config struct {
	RestHost string
	RestPort int
}
