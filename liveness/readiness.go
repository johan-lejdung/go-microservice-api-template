package liveness

import (
	"net/http"
)

func (a *API) InitReadinessRouter() {
	a.Router.
		Methods("GET").
		Path("/readiness/").
		Name("readiness").
		Handler(
			http.HandlerFunc(a.readiness()))
}

func (a *API) readiness() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
