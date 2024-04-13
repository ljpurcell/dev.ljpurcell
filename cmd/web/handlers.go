package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf("Could not parse template set in HOME handler: %v", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Printf("Could not write templated response in HOME handler: %v", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About me page..."))
}

func post(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	// validate slug

	fmt.Fprintf(w, "Show a specific blog post at slug: %v", slug)
}

func posts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List all the blog posts..."))
}

func projects(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List all my projects..."))
}
