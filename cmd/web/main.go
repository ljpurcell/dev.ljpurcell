package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/ljpurcell/dev.ljpurcell/internal/vcs"
	"golang.org/x/crypto/acme/autocert"
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
	EncodingExt string
	Nonce       string
	Post        Post
	Posts       map[string]*Post
}

type Post struct {
	Title    string
	Slug     string
	Category string
	Content  template.HTML
}

type contextKey string

const nonceKey contextKey = "nonce"
const encodingKey contextKey = "encoding"

const brEncoding string = "br"
const gzipEncoding string = "gzip"

var version string = vcs.Version()

func newHttpServer(addr string, handler http.Handler, logger *log.Logger) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      handler,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

}

func main() {

	displayVersion := flag.Bool("version", false, "Display version and exit")

	// Configuration
	addr := flag.String("addr", ":8080", "HTTP network address")
	staticDir := flag.String("staticDir", "./ui/static/", "Directory of the static assets")
	inProduction := flag.Bool("in-production", false, "Is the app runnning in a production environment")

	flag.Parse()

	if *displayVersion {
		fmt.Printf("Version:\t%s\n", version)
		os.Exit(0)
	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	// Formatters for code blocks
	// Class definitions for style in ./ui/static/chroma.css
	htmlBlockFormatter := html.New(html.WithClasses(true), html.TabWidth(4))
	if htmlBlockFormatter == nil {
		logger.Error("Could not create html block formatter")
		os.Exit(1)
	}

	htmlInlineFormatter := html.New(html.WithClasses(true), html.InlineCode(true))
	if htmlInlineFormatter == nil {
		logger.Error("Could not create html inline formatter")
		os.Exit(1)
	}

	// Syntax highlighting
	highlightStyle := styles.Get("github-dark")
	if highlightStyle == nil {
		logger.Error("Could not find style 'github-dark'")
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

	server := newHttpServer(
		*addr,
		app.routes(*staticDir),
		slog.NewLogLogger(logger.Handler(), slog.LevelError),
	)

	if *inProduction {
		certMan := &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Cache:      autocert.DirCache("tls"),
			HostPolicy: autocert.HostWhitelist("ljpurcell.com", "www.ljpurcell.com"),
		}

		server.TLSConfig = &tls.Config{
			GetCertificate: func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
				cert, err := certMan.GetCertificate(info)
				if err != nil {
					logger.Error("Failed to get TLS certificate", "error", err)
				}
				return cert, err
			},
			NextProtos: []string{"http/1.1", "h2"},
		}

		logger.Info("Starting TLS server...", "addr", server.Addr)

		go func() {
			err = server.ListenAndServeTLS("", "")
			logger.Error(err.Error())
			os.Exit(1)
		}()

		const httpPort = ":80"
		logger.Info("Starting HTTP redirect server...", "addr", httpPort)

		err = http.ListenAndServe(httpPort, certMan.HTTPHandler(nil))
		logger.Error(err.Error())
		os.Exit(1)
	} else {

		logger.Info("Starting HTTP server...", "addr", server.Addr)

		err = server.ListenAndServe()
		logger.Error(err.Error())
		os.Exit(1)
	}
}
