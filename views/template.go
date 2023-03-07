package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func handleParseResult(tpl *template.Template, err error) (Template, error) {
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

func Parse(filepath string) (Template, error) {
	return handleParseResult(template.ParseFiles(filepath))
}

func ParseFS(fs fs.FS, pattern string) (Template, error) {
	return handleParseResult(template.ParseFS(fs, pattern))
}

func must(tpl Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return tpl
}

func MustParse(filepath string) Template {
	return must(Parse(filepath))
}

func MustParseFS(fs fs.FS, pattern string) Template {
	return must(ParseFS(fs, pattern))
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
	}
}
