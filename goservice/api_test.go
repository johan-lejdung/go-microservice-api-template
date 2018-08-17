package goservice_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/johan-lejdung/go-microservice-api-template/goservice"
	"github.com/johan-lejdung/go-microservice-api-template/goservice/mocks"
	"github.com/stretchr/testify/assert"
)

//go:generate mockery -dir=../goservice -name=ServiceActions

func TestFail(t *testing.T) {
	goActions := &mocks.ServiceActions{}
	goActions.On("TestFunction").Return(errors.New("Error"))

	api := goservice.GoAPI{
		Router:    mux.NewRouter().StrictSlash(true),
		GoService: goActions,
	}
	api.InitAPIRoute()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/endpoint/", nil)

	api.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSuccess(t *testing.T) {
	goActions := &mocks.ServiceActions{}
	goActions.On("TestFunction").Return(nil)

	api := goservice.GoAPI{
		Router:    mux.NewRouter().StrictSlash(true),
		GoService: goActions,
	}
	api.InitAPIRoute()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/endpoint/", nil)

	api.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}
