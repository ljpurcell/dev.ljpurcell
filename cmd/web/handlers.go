package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About me page..."))
}

func (app *application) post(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	// validate slug

	fmt.Fprintf(w, "Show a specific blog post at slug: %v", slug)
}

func (app *application) posts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List all the blog posts..."))
}

func (app *application) projects(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List all my projects..."))
}

func (app *application) testMdPost(w http.ResponseWriter, r *http.Request) {
	html, err := app.renderMdFile(w, "./markdown/test.md")
	if err != nil {
		/*
		 * TODO: Check if file not found and return 404 if so
		 */
		app.serverError(w, r, err)
		return
	}

	w.Write(html)
}
