package main

import (
	"net/http"
	"net/url"
	"strings"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "home.tmpl.html", templateData{})
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "about.tmpl.html", templateData{})
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
	app.render(w, r, http.StatusOK, "posts.tmpl.html", templateData{
		Posts: app.postCache,
	})
}

func (app *application) project(w http.ResponseWriter, r *http.Request) {
	project := r.PathValue("project")

	var status int = http.StatusOK
	var page string

	switch project {
	case "portfolio":
		page = "portfolio.tmpl.html"
	case "sharks":
		page = "sharks.tmpl.html"
	case "tragics":
		page = "tragics.tmpl.html"
	case "prdy":
		page = "prdy.tmpl.html"
	case "got":
		page = "got.tmpl.html"
	case "fraudible":
		page = "fraudible.tmpl.html"
	default:
		status = http.StatusNotFound
		page = "not-found.tmpl.html" // TODO: Create page for 404 and 500
	}

	app.render(w, r, status, page, templateData{})
}

func (app *application) projects(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "projects.tmpl.html", templateData{})
}

func (app *application) tag(w http.ResponseWriter, r *http.Request) {
	tag := strings.Trim(r.PathValue("tag"), " ")

	posts, _ := app.tagCache[Tag(tag)]

	tagText, err := url.QueryUnescape(tag)
	if err != nil {
		app.serverError(w, r, err)
	}

	// Lowercases everything except the first letter of each word
	tagText = func(text string) string {
		words := strings.Fields(text)

		for i, word := range words {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}

		return strings.Join(words, " ")
	}(tagText)

	app.render(w, r, http.StatusOK, "tag_posts.tmpl.html", templateData{
		SelectedTag: Tag(tagText),
		TagPosts:    posts,
	})
}
