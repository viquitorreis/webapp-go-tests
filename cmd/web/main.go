package main

import (
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	Session *scs.SessionManager
}

func main() {
	// set up an app config
	app := &application{}

	// pegar o session manager
	app.Session = getSession()

	// print out a message
	log.Println("Starting server on 8080")

	// start the server
	if err := http.ListenAndServe(":8080", app.routes()); err != nil {
		log.Fatal(err)
	}
}
