package internal

import (
	"fmt"
	"log"
	"net/http"
)

var Handler *http.ServeMux

func logg(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Before")
			h.ServeHTTP(w, r) // call original
			fmt.Println("After")
		})
}

func StartServer() {
	k := logg(nil)
	fmt.Println(k)
	Handler := http.NewServeMux()

	// type HandlerFunc func(ResponseWriter, *Request)
	// internal mux call just casts our func to this ^
	// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	// 	if handler == nil {
	// 		panic("http: nil handler")
	// 	}
	// 	mux.Handle(pattern, HandlerFunc(handler))
	// }

	// type Handler interface {
	// 	ServeHTTP(ResponseWriter, *Request)
	// }

	server := &http.Server{
		Addr:    ":9000",
		Handler: Handler,
	}

	log.Println("Starting server on port 9000...")
	log.Fatalf("Server failed to start: %v", server.ListenAndServe())
}
