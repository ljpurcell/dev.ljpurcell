package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
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

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("Template %q does not exist", page)
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)

	var buf bytes.Buffer

	err := ts.ExecuteTemplate(&buf, "base", data)
	if err != nil {
		app.serverError(w, r, err)
	}

	buf.WriteTo(w)
}

func (app *application) generateNonce() (string, error) {
	nonce := make([]byte, 16)
	if _, err := rand.Read(nonce); err != nil {
		err = fmt.Errorf("Could not create nonce: %w", err)
		app.logger.Error(err.Error())
		return "", err
	}
	return base64.StdEncoding.EncodeToString(nonce), nil
}
