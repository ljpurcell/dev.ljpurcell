package main

import (
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/andybalholm/brotli"
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
		err := fmt.Errorf("template %q does not exist", page)
		app.serverError(w, r, err)
		return
	}

	data.Nonce = r.Context().Value(nonceKey).(string)
	encoding := r.Context().Value(encodingKey).(string)

	if encoding == brEncoding || encoding == gzipEncoding {
		data.EncodingExt = fmt.Sprintf(".%s", encoding)
	}

	var buf bytes.Buffer

	err := ts.ExecuteTemplate(&buf, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Vary", "Content-Encoding")

	if encoding == brEncoding {
		w.Header().Set("Content-Encoding", brEncoding)
		w.WriteHeader(status)

		bw := brotli.NewWriter(w)
		defer bw.Close()

		buf.WriteTo(bw)
		return
	} else if encoding == gzipEncoding {
		w.Header().Set("Content-Encoding", gzipEncoding)
		w.WriteHeader(status)

		zw := gzip.NewWriter(w)
		defer zw.Close()

		buf.WriteTo(zw)
		return
	}

	w.WriteHeader(status)
	buf.WriteTo(w)
}

func (app *application) generateNonce() (string, error) {
	nonce := make([]byte, 16)
	if _, err := rand.Read(nonce); err != nil {
		err = fmt.Errorf("could not create nonce: %w", err)
		app.logger.Error(err.Error())
		return "", err
	}
	return base64.StdEncoding.EncodeToString(nonce), nil
}
