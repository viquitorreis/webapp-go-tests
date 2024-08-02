package main

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

func getSession() *scs.SessionManager {
	// criando um novo session Manager
	session := scs.New()
	// tempo de vida da sessão
	session.Lifetime = 24 * time.Hour

	session.Cookie.Persist = true

	// isso vai evitar qualquer erro com versões atualizadas do browser
	session.Cookie.SameSite = http.SameSiteLaxMode

	// no localhost não vamos usar cookies encriptador mas em produção sim. JAMAIS USE cookies que não sejam encriptados
	session.Cookie.Secure = true

	return session
}
