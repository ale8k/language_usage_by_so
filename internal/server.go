package internal

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"os"

	"strconv"

	"github.com/ale8k/language_usage_by_so/internal/handlers"
)

var Server *http.Server

func StartServer() {
	// done := make(chan struct{})
	//go RenameMe(done)
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 9000
	}
	pprofEndpoints, err := strconv.ParseBool(os.Getenv("PPROF_ENDPOINTS"))
	if err != nil {
		log.Printf("Failed to set PPROF_ENDPOINTS")
		pprofEndpoints = false
	}

	mux := http.NewServeMux()

	if pprofEndpoints {
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	}

	Server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", "", strconv.Itoa(port)),
		Handler: mux,
	}

	handlers.RegisterAllHandlers(mux, Server)

	log.Printf("Starting server on port %v... \n", port)
	log.Fatalf("Server failed to start: %v", Server.ListenAndServe())
}
