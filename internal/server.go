package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ale8k/language_usage_by_so/internal/handlers"
)

var Server *http.Server

func StartServer() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 9000
	}

	mux := http.NewServeMux()

	Server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", "", strconv.Itoa(port)),
		Handler: mux,
	}

	handlers.RegisterAllHandlers(mux, Server)

	log.Printf("Starting server on port %v... \n", port)
	log.Fatalf("Server failed to start: %v", Server.ListenAndServe())
}
