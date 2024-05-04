package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

/*
Sets the headers for HTTP responses as well as static files.
*/
func (app *application) commonHeaders(next http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		nonce, err := app.generateNonce()
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		ctx := context.WithValue(r.Context(), nonceKey, nonce)

		acceptedEncodings := r.Header.Get("Accept-Encoding")

		if strings.Contains(acceptedEncodings, brEncoding) {
			w.Header().Set("Content-Encoding", brEncoding)
			ctx = context.WithValue(ctx, encodingKey, brEncoding)
		} else if strings.Contains(acceptedEncodings, gzipEncoding) {
			w.Header().Set("Content-Encoding", gzipEncoding)
			ctx = context.WithValue(ctx, encodingKey, gzipEncoding)
		}

		// Static files have a templateData field called EncodingExt that
		// lets them request the relevant encoded static file, so the URL.Path
		// will be something like /static/script.{js,css}.{br,gzip}
		if strings.Contains(r.URL.Path, ".js") {
			w.Header().Set("Content-Type", "application/javascript")
		} else if strings.Contains(r.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		}

		csp := fmt.Sprintf("default-src 'self'; style-src 'self' fonts.googleapis.com; font-src 'self' fonts.gstatic.com data:; script-src 'self' 'nonce-%s'", nonce)

		w.Header().Set("Content-Security-Policy", csp)
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
