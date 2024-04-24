package main

import (
	"errors"
	"net/http"
	"strings"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "home.tmpl.html", templateData{})
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

	// TODO: implement cache: cache[slug] => post

	path := strings.Join(pathBits, "")

	p := &Post{}
	if err := app.parseFileIntoPost(p, path); err != nil {
		if errors.Is(err, ErrPostNotFound) {
			app.clientError(w, 404)
			return
		}

		app.serverError(w, r, err)
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
