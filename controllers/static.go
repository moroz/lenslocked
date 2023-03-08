package controllers

import (
	"net/http"

	"github.com/moroz/lenslocked/templates"
	"github.com/moroz/lenslocked/views"
)

func StaticHandler(template views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		template.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := templates.GetFAQs()
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
