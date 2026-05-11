package main

import "net/http"

func (app *application) routes(staticDir string) http.Handler {
	mux := http.NewServeMux()

	// File server to serve static files
	fileServer := http.FileServer(safeFileSystem{http.Dir(staticDir)})
	mux.Handle("GET /static", http.NotFoundHandler())
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Application routes
	mux.HandleFunc("GET /{$}", app.home)

	return app.commonHeaders(mux)
}
