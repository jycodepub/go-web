package web

import (
	"html/template"
	"net/http"
)

type Template struct {
	tmpl *template.Template
}

func NewTemplate(name string, paths ...string) *Template {
	tmpl := template.Must(template.ParseFiles(paths...))
	return &Template{
		tmpl: tmpl,
	}
}

func (t *Template) Render(w http.ResponseWriter, context interface{}) {
	w.Header().Add("Content-Type", "text/html")
	if err := t.tmpl.Execute(w, context); err != nil {
		TextResponse(w, 500, "Template Error")
	}
}
