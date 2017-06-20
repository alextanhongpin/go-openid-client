package app

import (
	"html/template"
)

var templates map[string]*template.Template

func NewTemplate ()  map[string]*template.Template {
	if templates == nil {
		templates = make(map[string]*template.Template)

	}
	templates["home"] = template.Must(template.ParseFiles("templates/home.html"))
	return templates
}