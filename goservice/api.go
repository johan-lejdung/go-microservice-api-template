package goservice

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	log "github.com/sirupsen/logrus"
)

// GoAPI contains the configuration
type GoAPI struct {
	Router    *mux.Router    `inject:""`
	GoService ServiceActions `inject:""`
}

// InitAPIRoute initializes the route
func (a *GoAPI) InitAPIRoute() {

	a.Router.
		Methods("GET").
		Path("/endpoint/").
		Handler(negroni.New(
			negroni.HandlerFunc(a.GoEndpoint()),
		))
}

// GoEndpoint is an endpoint you should probably replace
func (a *GoAPI) GoEndpoint() func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		err := a.GoService.TestFunction()
		if err != nil {
			log.Debugf("Failed to execute test function. %s", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// return results
		w.WriteHeader(http.StatusOK)
	}
}
