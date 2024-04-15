package main

import (
	"net/http"
	"os"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	method := r.Method
	uri := r.URL.RequestURI()

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) renderMdFile(w http.ResponseWriter, file string) ([]byte, error) {
	contents, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var (
		b []byte
		p post
	)

	err = parseMdFile(contents, &p, &b)
	if err != nil {
		return nil, err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return mdToHTML(b), nil

}
