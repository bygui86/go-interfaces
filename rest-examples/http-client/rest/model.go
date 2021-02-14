package rest

import (
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

type Server struct {
	config     *Config
	router     *mux.Router
	httpServer *http.Server
	baseURL    *url.URL
	restClient HTTPClient
	running    bool
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Config struct {
	RestServerHost string
	RestServerPort int
	RestHost       string
	RestPort       int
}
