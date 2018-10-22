package goservice

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	log "github.com/sirupsen/logrus"
)

// GoAPI contains the configuration
type GoAPI struct {
	Router    *mux.Router `inject:""`
	GoService Services    `inject:""`
}

// InitAPIRoute initializes the route
func (a *GoAPI) InitAPIRoute() {
	a.Router.
		Methods("GET").
		Path("/endpoint/{id}").
		Handler(negroni.New(
			negroni.HandlerFunc(a.GoGetEndpoint()),
		))

	a.Router.
		Methods("POST").
		Path("/endpoint/").
		Handler(negroni.New(
			negroni.HandlerFunc(a.GoPostEndpoint()),
		))
}

// GoGetEndpoint is an endpoint you should probably replace
func (a *GoAPI) GoGetEndpoint() func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		vars := mux.Vars(r)
		// Get current project from Request
		ID, err := strconv.Atoi(vars["id"])
		if err != nil || ID == 0 {
			log.Debugf("Not using correct endpoint, ID not found or readable")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		value, err := a.GoService.GetFunction(ID)
		if err != nil {
			log.Debugf("Failed to execute get function. %s", err.Error())
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// return results
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(value); err != nil {
			panic(err)
		}
	}
}

// GoPostEndpoint is an endpoint you should probably replace
func (a *GoAPI) GoPostEndpoint() func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		// Get the payload
		payload := TestPayload{}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			log.Debugf("Failed to parse payload. %s", err.Error())
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		_, err := a.GoService.PostFunction(payload)
		if err != nil {
			log.Debugf("Failed to execute post function. %s", err.Error())
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// return results
		w.WriteHeader(http.StatusOK)
	}
}
