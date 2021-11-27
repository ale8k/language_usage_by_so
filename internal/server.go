package internal

import (
	"log"
	"net/http"
)

func StartServer() {
	server := &http.Server{
		Addr:    ":9000",
		Handler: nil,
	}
	log.Println("Starting server on port 9000...")
	log.Fatalf("Server failed to start: %v", server.ListenAndServe())
}
