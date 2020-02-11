package main

import (
	"net/http"
)

//create reqd middleware fn:
func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")

		//pass the flow to next
		next.ServeHTTP(w, r)
	})
}
