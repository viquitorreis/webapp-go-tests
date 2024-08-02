package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Has(t *testing.T) {
	form := NewForm(nil)
	has := form.Has("whatever")
	if has {
		t.Error("Form should not have field 'whatever'")
	}

	// aqui estamos criando uam form
	postedData := url.Values{}
	// key -> name, value -> john
	postedData.Add("name", "john")
	form = NewForm(postedData)

	has = form.Has("name")
	if !has {
		t.Error("Form should have field 'name'")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	// criando a form
	form := NewForm(r.PostForm)

	form.Required("name", "email", "password")

	if form.Valid() {
		t.Error("form should not be valid because required fields are empty")
	}

	postedDate := url.Values{}
	postedDate.Add("name", "john")
	postedDate.Add("email", "john@sample.com")
	postedDate.Add("password", "password")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedDate

	form = NewForm(r.PostForm)
	form.Required("name", "email", "password")

	// a esse ponto a form deveria ser válida
	if !form.Valid() {
		t.Error("form should be valid")
	}
}

func TestForm_Check(t *testing.T) {
	form := NewForm(nil)

	// nesse caso o form não é válido pois o campo password está vazio
	form.Check(false, "password", "password is required")
	if form.Valid() {
		t.Error("Valid() returns false, and it should be true when calling Check()")
	}
}

func TestForm_ErrorGet(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password is required")
	s := form.Errors.Get("password")

	if len(s) == 0 {
		t.Error("should have an error returned from Get, but does not")
	}

	s = form.Errors.Get("email")
	if len(s) != 0 {
		t.Error("should not have an error returned from Get, but does")
	}
}
