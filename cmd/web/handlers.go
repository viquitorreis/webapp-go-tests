package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

var pathToTemplates = "./templates/"

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	// passando dados para o template
	var td = make(map[string]any)

	if app.Session.Exists(r.Context(), "test") {
		// pegando o que tem na session
		msg := app.Session.GetString(r.Context(), "test")
		// passando para o template
		td["test"] = msg
	} else {
		app.Session.Put(r.Context(), "test", "Hit this page at"+time.Now().UTC().String())
	}

	err := app.render(w, r, "home.page.gohtml", &TemplateData{Data: td})
	if err != nil {
		fmt.Println(err)
	}
}

type TemplateData struct {
	IP   string // person ip
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, data *TemplateData) error {
	// parse template from disk
	parsedTemplate, err := template.ParseFiles(
		path.Join(pathToTemplates, name),
		path.Join(pathToTemplates, "base.layout.gohtml"),
	)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}

	data.IP = app.ipFromContext(r.Context())

	// execute template, passing data if any
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}

	return nil
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	// pegando a reqquest num formato que vamos conseguir trabalhar com a form
	err := r.ParseForm()
	if err != nil {
		log.Println("Erro ao passar a form", err)
		http.Error(w, "Erro", http.StatusBadRequest)
		return
	}

	// validando os dados
	form := NewForm(r.PostForm)
	form.Required("email", "password")

	if !form.Valid() {
		fmt.Fprintf(w, "Formulário inválido. Falhou a validação")
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	log.Println("Email: ", email)
	log.Println("Password: ", password)

	fmt.Fprint(w, email)
}
