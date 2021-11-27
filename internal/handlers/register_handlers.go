package handlers

import (
	"net/http"

	"github.com/ale8k/language_usage_by_so/internal/middleware"
)

// Registers all handlers to a route
// Each is manually included for explicit clarity
func RegisterAllHandlers(mux *http.ServeMux, serverInstance *http.Server) {
	// Health check endpoint
	mux.HandleFunc("/healthz", middleware.SetDefaultHeaders(getServerHealth, serverInstance.Addr))
}
