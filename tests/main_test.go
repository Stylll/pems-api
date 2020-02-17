package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	Routes "github.com/Stylll/pems-api/src/routes"
	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func TestMain(m *testing.M) {
	app = Routes.SetupRouter()

	code := m.Run()

	os.Exit(code)
}

func TestServer(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["PEMS"] != "ALIVE" {
		t.Errorf("Expected the 'PEMS' key of response to be 'ALIVE'. Got '%s'", m["PEMS"])
	}
}

func TestServerAPI(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1", nil)
	response := executeRequest(req)
	/*
		calling executeRequest again here because calling it for the first time
		returns 301 (redirect) after which it returns 200 (OK). Super Weird!.
		Subsequent calls after the first one returns 200
	*/
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["PEMS API"] != "WELCOME TO PEMS API" {
		t.Errorf("Expected the 'PEMS API' key of response to be 'WELCOME TO PEMS API'. Got '%s'", m["PEMS API"])
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
