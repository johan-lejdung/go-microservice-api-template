package utils_test

import (
	"testing"

	"github.com/johan-lejdung/go-microservice-api-template/utils"
	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	assert.EqualValues(t, utils.Min(1, 5), 1)
	assert.EqualValues(t, utils.Min(55, 5), 5)
	assert.EqualValues(t, utils.Min(10, 52), 10)
}

func TestMax(t *testing.T) {
	assert.EqualValues(t, utils.Max(1, 5), 5)
	assert.EqualValues(t, utils.Max(55, 5), 55)
	assert.EqualValues(t, utils.Max(10, 52), 52)
}

func TestMd5Hash(t *testing.T) {
	assert.Equal(t, utils.Md5Hash([]byte("5")), "e4da3b7fbbce2345d7772b0674a318d5")
	assert.Equal(t, utils.Md5Hash([]byte("testing")), "ae2b1fca515949e5d54fb22b8ed95575")
	assert.Equal(t, utils.Md5Hash([]byte("heellooa2")), "2eb2e29b8f83b3dc960c5068f32638d8")
}

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
		res := utils.StringInSlice(tt.value, slice)
		if res != tt.exceptedResult {
			t.Errorf("Sprintf(%s, slice) => %t, want %t", tt.value, res, tt.exceptedResult)
		}
	}
}
