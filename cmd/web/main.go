package main

import (
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
)

// To hold application dependencies, enabling dependency injection
type application struct {
	logger *slog.Logger
}

type post struct {
	title     string
	slug      string
	catergory string
	content   template.HTML
}

func main() {

	// Configuration
	addr := flag.String("addr", ":8080", "HTTP network address")
	staticDir := flag.String("staticDir", "./ui/static/", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()

	// File server to serve static files
	fileServer := http.FileServer(safeFileSystem{http.Dir(*staticDir)})
	mux.Handle("GET /static", http.NotFoundHandler())
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Application routes
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /about", app.about)
	mux.HandleFunc("GET /post/{slug}", app.post)
	mux.HandleFunc("GET /posts", app.posts)
	mux.HandleFunc("GET /projects", app.projects)

	// TEST ROUTES
	mux.HandleFunc("GET /test", app.testMdPost)

	logger.Info("Starting server...", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
