package main

import (
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/styles"
)

// To hold application dependencies, enabling dependency injection
type application struct {
	logger *slog.Logger

	// Syntax highlighting
	defaultLang         string
	htmlBlockFormatter  *html.Formatter
	htmlInlineFormatter *html.Formatter
	highlightStyle      *chroma.Style
}

type Post struct {
	Title     string
	Slug      string
	Catergory string
	Content   template.HTML
}

func main() {

	// Configuration
	addr := flag.String("addr", ":8080", "HTTP network address")
	staticDir := flag.String("staticDir", "./ui/static/", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	htmlBlockFormatter := html.New(html.WithClasses(false), html.TabWidth(4))
	if htmlBlockFormatter == nil {
		logger.Error("Could not create html block formatter")
		os.Exit(1)
	}

	htmlInlineFormatter := html.New(html.WithClasses(false), html.InlineCode(true))

	if htmlInlineFormatter == nil {
		logger.Error("Could not create html inline formatter")
		os.Exit(1)
	}

	styleName := styles.GitHubDark.Name
	highlightStyle := styles.Get(styleName)
	if highlightStyle == nil {
		logger.Error(fmt.Sprintf("Could not find style %s", styleName))
		os.Exit(1)
	}

	app := &application{
		defaultLang:         "go",
		logger:              logger,
		htmlBlockFormatter:  htmlBlockFormatter,
		htmlInlineFormatter: htmlInlineFormatter,
		highlightStyle:      highlightStyle,
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
