package main

import (
	"log"
	"net/http"
)

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
