package handlers

import (
	"encoding/json"
	"mime"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var HealthzCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: "server",
		Subsystem: "health",
		Name:      "healthz_requests_total",
		Help:      "The total amount of requests sent to healthz",
	},
	[]string{"caller"},
)

func getServerHealth(rw http.ResponseWriter, req *http.Request) {
	HealthzCounter.WithLabelValues(req.RemoteAddr).Inc()
	data, _ := json.Marshal(struct {
		Healthy bool `json:"healthy"`
	}{true})

	rw.Header().Add("Content-Type", mime.TypeByExtension(".json"))
	rw.Write(data)
}
