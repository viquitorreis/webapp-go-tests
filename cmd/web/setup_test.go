package main

import (
	"os"
	"testing"
)

// por mais que seja uma variável a nível de pacote, quando estamos executando os testes
// ou fazendo o build da aplicação, os arquivos de teste são ignorados. São apenas usados
// quando estamos rodando os testes
var app application

// IMPORTANTE! Essa função vai sempre rodar antes de qualquer outro teste
func TestMain(m *testing.M) {
	// podemos fazer todos os testes de session e database aqui
	app.Session = getSession()

	os.Exit(m.Run())
}
