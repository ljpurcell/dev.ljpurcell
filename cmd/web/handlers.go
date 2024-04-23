package main

import (
	"html/template"
	"net/http"
	"strings"
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

	// TODO: validate slug
	pathBits := []string{
		"./markdown/",
		slug,
		".md",
	}

	path := strings.Join(pathBits, "")

	p := &Post{}
	if err := app.parseFileIntoPost(p, path); err != nil {
		/*
		 * TODO: Check if file not found and return 404 if so
		 */
		app.serverError(w, r, err)
		return
	}

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/post.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = ts.ExecuteTemplate(w, "base", p)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) posts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List all the blog posts..."))
}

func (app *application) projects(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List all my projects..."))
}
