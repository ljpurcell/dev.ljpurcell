package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	// Configuration
	addr := flag.String("addr", ":8080", "HTTP network address")
	staticDir := flag.String("staticDir", "./ui/static/", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux := http.NewServeMux()

	// File server to serve static files
	fileServer := http.FileServer(safeFileSystem{http.Dir(*staticDir)})
	mux.Handle("GET /static", http.NotFoundHandler())
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Application routes
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /about", about)
	mux.HandleFunc("GET /post/{slug}", post)
	mux.HandleFunc("GET /posts", posts)
	mux.HandleFunc("GET /projects", projects)

	log.Printf("Starting server on %q\n", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
