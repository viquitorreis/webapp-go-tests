package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_application_routes(t *testing.T) {
	var registered = []struct {
		route  string
		method string
	}{
		{"/", "GET"},
		{"/static/*", "GET"},
	}

	var app application
	mux := app.routes()

	// chiRoutes := mux.(*chi.Mux).Routes()
	chiRoutes := mux.(chi.Routes)

	for _, r := range registered {
		if !routeExists(r.route, r.method, chiRoutes) {
			t.Errorf("route %s %s not found", r.method, r.route)
		}
	}
}

func routeExists(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	found := false

	_ = chi.Walk(
		chiRoutes,
		func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
			if strings.EqualFold(method, testMethod) && strings.EqualFold(route, testRoute) {
				found = true
			}
			return nil // continue walking the router
		},
	)

	return found
}
