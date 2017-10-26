package testutil

import (
	"net/http"
	"reflect"
	"testing"
)

func ShouldBe(t *testing.T, name string, shouldValue interface{}, isValue interface{}) {
	t.Errorf("%v should be %v but is %v", name, shouldValue, isValue)
}

func ShouldBeEqual(t *testing.T, name string, shouldValue interface{}, isValue interface{}) {
	if !reflect.DeepEqual(isValue, shouldValue) {
		ShouldBe(t, name, shouldValue, isValue)
	}
}

func HTTPError(t *testing.T, url string, err interface{}) {
	t.Errorf("%v could not be reached: %v", url, err)
}

func TestGet(t *testing.T, url string) *http.Response {
	resp, err := http.Get("http://" + url)
	if err != nil {
		HTTPError(t, url, err)
	}

	return resp
}

func Error(t *testing.T, err error) {
	t.Errorf("function under test raise error %v", err)
}

func CheckError(t *testing.T, err error) {
	if err != nil {
		Error(t, err)
	}
}
