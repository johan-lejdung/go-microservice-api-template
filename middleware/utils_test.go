package middleware_test

import (
	"testing"

	"github.com/johan-lejdung/go-microservice-api-template/middleware"
)

var stringInSliceTest = []struct {
	value          string
	exceptedResult bool
}{
	{"inslice", true},
	{"alsoinslice", true},
	{"test0002", true},
	{"asd", false},
	{"notinslice", false},
	{"adfggg", false},
}

func TestStringInSlice(t *testing.T) {
	slice := []string{"inslice", "alsoinslice", "test0002", "test"}
	for _, tt := range stringInSliceTest {
		res := middleware.StringInSlice(tt.value, slice)
		if res != tt.exceptedResult {
			t.Errorf("Sprintf(%s, slice) => %t, want %t", tt.value, res, tt.exceptedResult)
		}
	}
}
