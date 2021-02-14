package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"

	"github.com/bygui86/go-testing/rest-examples/http-client/commons"
	"github.com/bygui86/go-testing/rest-examples/http-client/logging"
)

const (
	// urls
	rootProductsEndpoint     = "/products"
	productsIdEndpoint       = rootProductsEndpoint + "/{id:[0-9]+}" // used by mux router
	productsIdServerEndpoint = rootProductsEndpoint + "/%d"          // used to reach http-server

	// headers
	// keys
	headerContentType  = "Content-Type"
	headerAccept       = "Accept"
	headerUserAgent    = "User-Agent"
	headerCustomSource = "Custom-Source"
	// values
	headerUserAgentClient = "GoTracesHttpClient/1.0"
	headerApplicationJson = "application/json"
)

// SERVER

func (s *Server) setupRouter() {
	logging.Log.Debug("Create new router")

	s.router = mux.NewRouter().StrictSlash(true)

	s.router.Use(requestInfoPrintingMiddleware)

	s.router.HandleFunc(rootProductsEndpoint, s.GetProducts).Methods(http.MethodGet)
	s.router.HandleFunc(productsIdEndpoint, s.GetProduct).Methods(http.MethodGet)
	s.router.HandleFunc(rootProductsEndpoint, s.CreateProduct).Methods(http.MethodPost)
	s.router.HandleFunc(productsIdEndpoint, s.UpdateProduct).Methods(http.MethodPut)
	s.router.HandleFunc(productsIdEndpoint, s.DeleteProduct).Methods(http.MethodDelete)
}

func (s *Server) setupHTTPServer() {
	logging.SugaredLog.Debugf("Create new HTTP server on port %d", s.config.RestPort)

	if s.config != nil {
		s.httpServer = &http.Server{
			Addr:    fmt.Sprintf(commons.HttpServerHostFormat, s.config.RestHost, s.config.RestPort),
			Handler: s.router,
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: commons.HttpServerWriteTimeoutDefault,
			ReadTimeout:  commons.HttpServerReadTimeoutDefault,
			IdleTimeout:  commons.HttpServerIdelTimeoutDefault,
		}
		return
	}

	logging.Log.Error("HTTP server creation failed: REST server configurations not loaded")
}

// HANDLERS

func sendJsonResponse(writer http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	writer.Header().Set(headerContentType, headerApplicationJson)
	writer.WriteHeader(code)
	_, writeErr := writer.Write(response)
	if writeErr != nil {
		logging.SugaredLog.Errorf("Error sending JSON response: %s", writeErr.Error())
	}
}

func sendErrorResponse(writer http.ResponseWriter, code int, message string) {
	sendJsonResponse(writer, code, map[string]string{"error": message})
}

func closeBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		logging.SugaredLog.Errorf("Closing request body failed: %s", err.Error())
	}
}

// OTHERS

func CreateBaseUrl(host string, port int) (*url.URL, error) {
	return url.Parse(fmt.Sprintf("http://%s:%d", host, port))
}

func CreateRestClient() *http.Client {
	logging.Log.Debug("Create new HTTP client")

	return &http.Client{
		Timeout: 3 * time.Second,
	}
}

func (s *Server) setRequestHeaders(restRequest *http.Request) {
	restRequest.Header.Set(headerAccept, headerApplicationJson)
	restRequest.Header.Set(headerUserAgent, headerUserAgentClient)
}
