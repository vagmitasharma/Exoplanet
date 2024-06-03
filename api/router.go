package api

import (
	_ "embed"

	"github.com/gorilla/mux"
)

// ScopeRouter is a standard way for feature scopes to setup routing.
type ScopeRouter interface {
	Route(base, api *mux.Router)
}

// newRouter sets up middleware and routing.
func newRouter() (*mux.Router, *mux.Router) {
	baseRouter := mux.NewRouter().StrictSlash(true)

	apiRouter := baseRouter.PathPrefix("/exoplanets").Subrouter()

	return baseRouter, apiRouter

}
