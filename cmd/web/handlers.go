package main

import (
	"fmt"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! This is the home page.")
}

type TemplateData struct {
	IP string // person ip
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, data *TemplateData) error {
	return nil
}
