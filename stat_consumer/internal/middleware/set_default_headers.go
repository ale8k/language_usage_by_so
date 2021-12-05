package middleware

import (
	"net/http"
)

// Adds base headers to all requests of language_usage_server
func SetDefaultHeaders(handlerFunc http.HandlerFunc, addr string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		h := rw.Header()
		h.Add("Server", "GoLang net/http") // what is appropriate for this in golang?
		h.Add("Access-Control-Allow-Origin", "*")
		h.Add("Connection", "keep-alive")
		h.Add("Keep-Alive", "timeout=3, max=1000")
		h.Add("X-Server-Address", addr)
		handlerFunc.ServeHTTP(rw, r)
	}
}
