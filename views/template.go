package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
)

type Template struct {
	htmlTpl *template.Template
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	tpl := template.New(patterns[0])
	// Placeholder implementation so that parsing does not panic
	tpl = tpl.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return ""
		},
	})
	tpl, err := tpl.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

func must(tpl Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return tpl
}

func MustParseFS(fs fs.FS, patterns ...string) Template {
	return must(ParseFS(fs, patterns...))
}

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}) {
	tpl := t.htmlTpl
	// actual implementation
	tpl = tpl.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrf.TemplateField(r)
		},
	})
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tpl.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
	}
}
