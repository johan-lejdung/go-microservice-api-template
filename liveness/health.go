package liveness

import (
	"net/http"
)

func (a *API) InitHealthRouter() {
	a.Router.
		Methods("GET").
		Path("/health/").
		Name("health").
		Handler(http.HandlerFunc(a.health()))

	a.Router.
		Methods("GET").
		Path("/").
		Name("health").
		Handler(http.HandlerFunc(a.health()))
}

func (a *API) health() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
