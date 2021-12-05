package handlers

import "github.com/prometheus/client_golang/prometheus"

func init() {
	initialiseAllMetrics()
}

func initialiseAllMetrics() {
	prometheus.MustRegister(HealthzCounter)
}
