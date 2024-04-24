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

	postCache     map[string]*Post
	templateCache map[string]*template.Template
}

type templateData struct {
	Post Post
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

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	// Formatters for code blocks
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

	// Syntax highlighting
	styleName := styles.GitHubDark.Name
	highlightStyle := styles.Get(styleName)
	if highlightStyle == nil {
		logger.Error(fmt.Sprintf("Could not find style %s", styleName))
		os.Exit(1)
	}

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(fmt.Sprintf("Could not create template cache: %v", err))
		os.Exit(1)
	}

	app := &application{
		defaultLang:         "go",
		logger:              logger,
		htmlBlockFormatter:  htmlBlockFormatter,
		htmlInlineFormatter: htmlInlineFormatter,
		highlightStyle:      highlightStyle,
		templateCache:       templateCache,
	}

	postCache, err := app.newPostCache()
	if err != nil {
		logger.Error(fmt.Sprintf("Could not create post cache: %v", err))
		os.Exit(1)
	}

	app.postCache = postCache

	logger.Info("Starting server...", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes(*staticDir))
	logger.Error(err.Error())
	os.Exit(1)
}
