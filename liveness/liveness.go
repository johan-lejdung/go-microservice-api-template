package liveness

import "github.com/gorilla/mux"

type LivenessAPI interface {
	InitHealthRouter()
	InitReadinessRouter()
}

type API struct {
	Router *mux.Router `inject:""`
}

var _ LivenessAPI = &API{}
