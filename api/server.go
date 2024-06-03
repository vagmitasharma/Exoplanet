package api

import (
	"context"
	_ "embed"
	"exoplanet/config"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
)

// Serve will setup all scope routing, and listen for HTTP requests.
func Serve(ctx context.Context, scopes []ScopeRouter, errChan chan<- error, c config.Config) {
	baseRouter, apiRouter := newRouter()

	// setup routing for each scope
	for _, scope := range scopes {
		scope.Route(baseRouter, apiRouter)
	}

	// CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"DELETE", "GET", "HEAD", "POST", "PUT", "PATCH", "OPTIONS"})
	handler := handlers.CORS(originsOk, headersOk, methodsOk)(baseRouter)

	cfg := getServerConfig(c)
	log.Println("HTTP server start listening")
	httpServ := &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf(":%v", cfg.Port),
		WriteTimeout: cfg.Timeout * time.Second,
		ReadTimeout:  cfg.Timeout * time.Second,
	}
	go func() {
		errChan <- httpServ.ListenAndServe()
	}()
}

type serverConfig struct {
	Port    string
	Timeout time.Duration
}

func getServerConfig(config config.Config) serverConfig {
	c := serverConfig{}
	c.Port = config.PORT
	c.Timeout = 5
	if config.ENVIRONMENT == "dev" {
		c.Timeout = 500 // extend timeout to make development easier
	}
	return c
}
