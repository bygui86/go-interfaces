package rest

import (
	"context"
	"fmt"
	"time"

	"github.com/bygui86/go-testing/rest-examples/http-server/database"
	"github.com/bygui86/go-testing/rest-examples/http-server/logging"
)

func New(cfg *Config, db database.InMemoryDb) *Server {
	logging.Log.Info("Create new REST server")

	server := &Server{
		config: cfg,
		db:     db,
	}

	server.setupRouter()
	server.setupHTTPServer()
	return server
}

func (s *Server) Start() error {
	logging.Log.Info("Start REST server")

	if s.httpServer != nil && !s.running {
		var err error
		go func() {
			err = s.httpServer.ListenAndServe()
			if err != nil {
				logging.SugaredLog.Errorf("REST server start failed: %s", err.Error())
			}
		}()
		if err != nil {
			return err
		}
		s.running = true
		logging.SugaredLog.Infof("REST server listening on port %d", s.config.RestPort)
		return nil
	}

	return fmt.Errorf("REST server start failed: HTTP server not initialized or HTTP server already running")
}

func (s *Server) Shutdown(timeout int) {
	logging.SugaredLog.Warnf("Shutdown REST server, timeout %d", timeout)

	if s.httpServer != nil && s.running {
		// create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
		// does not block if no connections, otherwise wait until the timeout deadline
		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			logging.SugaredLog.Errorf("Error shutting down REST server: %s", err.Error())
		}

		s.running = false
		return
	}

	logging.Log.Error("REST server shutdown failed: HTTP server not initialized or HTTP server not running")
}
