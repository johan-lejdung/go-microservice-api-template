package goservice_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/johan-lejdung/go-microservice-api-template/goservice"
	"github.com/johan-lejdung/go-microservice-api-template/goservice/mocks"
	"github.com/stretchr/testify/assert"
)

//go:generate mockery -dir=../goservice -name=Services

func TestGET_WrongID(t *testing.T) {
	goActions := &mocks.Services{}
	goActions.On("GetFunction", 1).Return("", errors.New("Error"))

	api := goservice.GoAPI{
		Router:    mux.NewRouter().StrictSlash(true),
		GoService: goActions,
	}
	api.InitAPIRoute()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/endpoint/asd", nil)

	api.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGET_Fail(t *testing.T) {
	goActions := &mocks.Services{}
	goActions.On("GetFunction", 1).Return("", errors.New("Error"))

	api := goservice.GoAPI{
		Router:    mux.NewRouter().StrictSlash(true),
		GoService: goActions,
	}
	api.InitAPIRoute()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/endpoint/1", nil)

	api.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGET_Success(t *testing.T) {
	goActions := &mocks.Services{}
	goActions.On("GetFunction", 1).Return("value", nil)

	api := goservice.GoAPI{
		Router:    mux.NewRouter().StrictSlash(true),
		GoService: goActions,
	}
	api.InitAPIRoute()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/endpoint/1", nil)

	api.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPOST_Fail(t *testing.T) {
	payload := goservice.TestPayload{
		Value: "This is a value",
	}
	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	goActions := &mocks.Services{}
	goActions.On("PostFunction", payload).Return(0, errors.New("Error"))

	api := goservice.GoAPI{
		Router:    mux.NewRouter().StrictSlash(true),
		GoService: goActions,
	}
	api.InitAPIRoute()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/endpoint/", bytes.NewBuffer(body))

	api.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPOST_Success(t *testing.T) {
	payload := goservice.TestPayload{
		Value: "This is a value",
	}
	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	goActions := &mocks.Services{}
	goActions.On("PostFunction", payload).Return(1, nil)

	api := goservice.GoAPI{
		Router:    mux.NewRouter().StrictSlash(true),
		GoService: goActions,
	}
	api.InitAPIRoute()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/endpoint/", bytes.NewBuffer(body))

	api.Router.ServeHTTP(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}
