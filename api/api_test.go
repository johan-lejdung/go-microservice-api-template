package api_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/johan-lejdung/go-microservice-api-template/api/mocks"

	"github.com/johan-lejdung/go-microservice-api-template/api"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

//go:generate mockery -dir=../goservice -name=ServiceActions

func TestFail(t *testing.T) {
	goActions := &mocks.ServiceActions{}
	goActions.On("TestFunction").Return(errors.New("Error"))

	api := api.GoAPI{
		Router:    mux.NewRouter().StrictSlash(true),
		GoService: goActions,
	}
	api.InitAPIRoute()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/cron/generate-birthday-reminders/", nil)

	api.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSuccess(t *testing.T) {
	goActions := &mocks.ServiceActions{}
	goActions.On("TestFunction").Return(nil)

	api := api.GoAPI{
		Router:    mux.NewRouter().StrictSlash(true),
		GoService: goActions,
	}
	api.InitAPIRoute()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/cron/generate-birthday-reminders/", nil)

	api.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}
