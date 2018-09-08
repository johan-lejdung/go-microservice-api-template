package liveness_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/johan-lejdung/go-microservice-api-template/liveness"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	api := &liveness.API{
		Router: mux.NewRouter().StrictSlash(true),
	}
	api.InitHealthRouter()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/health/", nil)
	api.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}
