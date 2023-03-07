package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/moroz/lenslocked/controllers"
	"github.com/moroz/lenslocked/templates"
	"github.com/moroz/lenslocked/views"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page not found", http.StatusNotFound)
}

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(
		views.MustParseFS(templates.FS, "home.gohtml")))
	r.Get("/contact", controllers.StaticHandler(
		views.MustParseFS(templates.FS, "contact.gohtml")))
	r.Get("/faq", controllers.StaticHandler(
		views.MustParseFS(templates.FS, "faq.gohtml")))
	r.NotFound(notFoundHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
