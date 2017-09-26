package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Init()

	code := m.Run()

	os.Exit(code)
}

// Test hello api without Router.Run()
func TestHelloApi(t *testing.T) {
	req, _ := http.NewRequest("GET", "/hello/Foo", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "{\"msg\":\"Hello Bar\"}" {
		t.Errorf("Expected '{\"msg\":\"Hello Bar\"}'. Got '%s'.", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
