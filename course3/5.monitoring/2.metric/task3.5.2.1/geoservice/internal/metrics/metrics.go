package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	LoginCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "Login_Counter",
		Help: "Total number of logins",
	})
	RegisterCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "Register_Counter",
		Help: "Total number of registrations",
	})
	LoginDurationCounter = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "login_duration_seconds",
		Help:    "Login duration in seconds",
		Buckets: prometheus.LinearBuckets(0.1, 0.1, 10),
	})
	RegisterDurationCounter = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "Register_duration_seconds",
		Help:    "Register duration in seconds",
		Buckets: prometheus.LinearBuckets(0.1, 0.1, 10),
	})
)
