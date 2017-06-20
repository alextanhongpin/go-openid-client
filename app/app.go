package app

import (
	"html/template"

	"github.com/julienschmidt/httprouter"
)

type Env struct {
	templates map[string]*template.Template
	router    *httprouter.Router
}

func New(t map[string]*template.Template, r *httprouter.Router) *Env {
	return &Env{
		templates: t,
		router:    r,
	}
}
