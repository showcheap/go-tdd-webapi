package main

import (
	"net/http"
	"os"
	"testing"

	"net/http/httptest"
)

var a App

func TestMain(m *testing.M) {
	a = App{}

	// log.Println("Init test")

	a.Initialize()

	code := m.Run()

	os.Exit(code)
}

func TestEmptyProduct(t *testing.T) {
	req, _ := http.NewRequest("GET", "/products", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expectd an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
