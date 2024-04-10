package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Home"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About me page..."))
}

func post(w http.ResponseWriter, r *http.Request) {
    slug := r.PathValue("slug")
    w.Write([]byte("Show a specific blog post at slug: " + slug))
}

func posts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List all the blog posts..."))
}

func projects(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List all my projects..."))
}

func main() {

    const port = ":4000"

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /about", about)
	mux.HandleFunc("GET /post/{slug}", post)
	mux.HandleFunc("GET /posts", posts)
	mux.HandleFunc("GET /projects", projects)

	log.Printf("Starting server on %q\n", port)

	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
