package controllers

import (
	"net/http"

	"github.com/moroz/lenslocked/views"
)

type Static struct {
	Template views.Template
}

func (static Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	static.Template.Execute(w, nil)
}
