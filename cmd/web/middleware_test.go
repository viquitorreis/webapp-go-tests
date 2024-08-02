package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_application_addIPToContext(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		expectedIP  string
		emptyIP     bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		// se os números forem menores que 255, o IP é válido
		{"X-Forwarded-For", "192.3.2.1", "", false},
		{"", "", "hello:world", false},
	}

	// create a dummy handler que vai ser usado para checar o contexto
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// precisamos checar se o IP foi adicionado ao contexto
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error("Context should have a value")
		}

		// checar se o valor é uma string
		ip, ok := val.(string) // type assertion
		if !ok {
			t.Error("Context value should be a string")
		}

		t.Log("IP from context: ", ip)
	})

	var app application

	for _, e := range tests {
		// vamos criar o handler para testar
		handlerToTest := app.addIPToContext(nextHandler)

		// criar uma mock request
		req := httptest.NewRequest("GET", "http://testing", nil)
		if e.emptyIP { // caso o endereço for vazio
			req.RemoteAddr = ""
		}

		// adicionando um header na request antes de chamar o handler para testar
		if len(e.headerName) > 0 {
			req.Header.Add(e.headerName, e.headerValue)
		}

		// caso exista o endereço
		if len(e.expectedIP) > 0 {
			req.RemoteAddr = e.expectedIP
		}

		// testando o handler
		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func Test_application_ipFromContext(t *testing.T) {
	// criar um app var do tipo application
	var app application

	// pegar o context
	ctx := context.Background()

	// colocar um valor no context
	ctx = context.WithValue(ctx, contextUserKey, "whatever.ip.user.have")

	// chamar a function
	ip := app.ipFromContext(ctx)

	// performar o teste
	if !strings.EqualFold(ip, "whatever.ip.user.have") {
		t.Errorf("Expected: %s, got: %s", "whatever.ip.user.have", ip)
	}
}
