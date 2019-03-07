package main

import (
	"bytes"
	"encoding/json"
	"go-api-tdd/models"
	"log"
	"net/http"
	"os"
	"testing"

	"net/http/httptest"
)

var a App

func TestMain(m *testing.M) {
	a = App{}

	// log.Println("Init test")

	a.Initialize("test.db")

	clearTable()

	code := m.Run()

	// log.Println("Test end")
	a.DB.Close()
	os.Exit(code)

}

func TestEmptyProduct(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/product", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expectd an empty array. Got %s", body)
	}
}

func TestGetNonExistProduct(t *testing.T) {
	req, _ := http.NewRequest("GET", "/product/11", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, res.Code)

	var m map[string]string
	json.Unmarshal(res.Body.Bytes(), &m)
	if m["error"] != "Product not found" {
		t.Errorf("Expected the 'error' key of the response 'Product not found'. Got %s", m["error"])
	}
}

func TestCreateProudct(t *testing.T) {
	clearTable()

	payload := []byte(`{"name":"test product", "price":11.22}`)

	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(payload))
	res := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, res.Code)

	var m map[string]interface{}
	json.Unmarshal(res.Body.Bytes(), &m)

	if m["name"] != "test product" {
		t.Errorf("Expected product name to 'test product'. Got %v", m["name"])
	}

	if m["price"] != 11.22 {
		t.Errorf("Expected product price to '11.22'. Got %v", m["price"])
	}

	if m["id"] != 1.0 {
		t.Errorf("Expected product id to '1'. Got %v", m["id"])
	}
}

func clearTable() {
	q := a.DB.Delete(models.Product{})
	a.DB.Exec("DELETE FROM sqlite_sequence")

	if q.Error != nil {
		log.Panic("Error clear table", q.Error.Error())
	}

	// log.Printf("Delete rows %v", q.RowsAffected)
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
