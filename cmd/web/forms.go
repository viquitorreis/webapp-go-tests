package main

import (
	"net/url"
	"strings"
)

// vamos armazenar os erros associados com a form.
// o indice vai ser uma string, e os valores um slice de strings pois podemos ter vários tipos de erros, pois pode falhar em mais de um tipo de erro ao mesmo tempo
type errors map[string][]string

func (e errors) Get(field string) string {
	// vamos tentar pegar o erro associado com o campo
	errorSlice := e[field]
	if len(errorSlice) == 0 {
		return ""
	}

	return errorSlice[0] // retornamos o primeiro erro
}

func (e errors) Add(field, message string) {
	// adicionamos o erro ao slice de erros associado com o campo
	e[field] = append(e[field], message)
}

type Form struct {
	// vamos armazenar os valores da form
	Data   url.Values
	Errors errors
}

func NewForm(date url.Values) *Form {
	return &Form{
		Data:   date,
		Errors: map[string][]string{},
	}
}

func (f *Form) Has(field string) bool {
	x := f.Data.Get(field)
	// could have done return x != ""
	if x == "" {
		return false
	}
	return true
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Data.Get(field)

		// se o campo estiver vazio, adicionamos um erro
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "Este campo não pode estar vazio")
		}
	}
}

func (f *Form) Check(ok bool, key, message string) {
	// key será o nome do campo na nossa form
	// message é qualquer tipo de mensagem de erro que recebermos como ultimo parametro para checar
	if !ok {
		f.Errors.Add(key, message)
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
