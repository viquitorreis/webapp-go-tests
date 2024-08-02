package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	var theTests = []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"home", "/", http.StatusOK},
		{"404", "/fish", http.StatusNotFound},
	}

	routes := app.routes()

	// create a new test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close() // close the server when test ends

	pathToTemplates = "./../../templates/"

	// range over the tests
	for _, tt := range theTests {
		resp, err := ts.Client().Get(ts.URL + tt.url)
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}

		if resp.StatusCode != tt.expectedStatusCode {
			t.Errorf("for %s: expected %d; got %d", tt.name, tt.expectedStatusCode, resp.StatusCode)
		}
	}
}
