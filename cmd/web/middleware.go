package main

import (
	"context"
	"fmt"
	"net/http"
)

func (app *application) commonHeaders(next http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		nonce, err := app.generateNonce()
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		ctx := context.WithValue(r.Context(), nonceKey, nonce)

		csp := fmt.Sprintf("default-src 'self'; style-src 'self' cdn.xeiaso.net; font-src 'self' data:; script-src 'self' 'nonce-%s'", nonce)

		w.Header().Set("Content-Security-Policy", csp)
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
