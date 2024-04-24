package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "home.tmpl.html", templateData{})
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About me page..."))
}

func (app *application) post(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	p, ok := app.postCache[slug]
	if !ok {
		app.clientError(w, 404)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	app.render(w, r, http.StatusOK, "post.tmpl.html", templateData{
		Post: *p,
	})
}

func (app *application) posts(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "posts.tmpl.html", templateData{})
}

func (app *application) projects(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "projects.tmpl.html", templateData{})
}
