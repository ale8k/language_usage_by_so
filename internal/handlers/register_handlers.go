package handlers

import (
	"net/http"

	"github.com/ale8k/language_usage_by_so/internal/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Registers all handlers to a route
// Each is manually included for explicit clarity
func RegisterAllHandlers(mux *http.ServeMux, serverInstance *http.Server) {
	// Register Prometheus metrics handler
	mux.Handle("/metrics", promhttp.Handler())
	// Health check endpoint
	mux.HandleFunc("/healthz", middleware.SetDefaultHeaders(getServerHealth, serverInstance.Addr))
}
