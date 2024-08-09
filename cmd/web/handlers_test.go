package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestAppHomeOld(t *testing.T) {
	// criar a request
	req, _ := http.NewRequest("GET", "/", nil)

	req = addContextAndSessionToRequest(req, app)

	// criando um response writer
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(app.Home)

	handler.ServeHTTP(rr, req)

	// checando o status code
	if rr.Code != http.StatusOK {
		t.Errorf("TestAppHome return wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	body, _ := io.ReadAll(rr.Body)
	want := string("<small>From Session:")
	if !strings.Contains(string(body), want) {
		t.Errorf("TestAppHome: body does not contain %s", want)
	}
}

func getCtx(req *http.Request) context.Context {
	ctx := context.WithValue(req.Context(), contextUserKey, "unknown")
	return ctx
}

func addContextAndSessionToRequest(req *http.Request, app application) *http.Request {
	req = req.WithContext(getCtx(req))
	// o header X-Session é esperado que esteja presente na request se formos usar a session
	ctx, _ := app.Session.Load(req.Context(), req.Header.Get("X-Session"))

	return req.WithContext(ctx)
}
