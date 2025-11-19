package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	reqCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "example_requests_total",
			Help: "Total number of requests received",
		},
		[]string{"path", "code"},
	)

	reqDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "example_request_duration_seconds",
			Help:    "Request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
	)
)

func init() {
	prometheus.MustRegister(reqCounter)
	prometheus.MustRegister(reqDuration)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	sleep := time.Duration(rand.Intn(200)) * time.Millisecond
	time.Sleep(sleep)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello from example service!")

	dur := time.Since(start).Seconds()
	reqDuration.Observe(dur)
	reqCounter.WithLabelValues(r.URL.Path, "200").Inc()

	log.Printf("msg=handled_request path=%s duration=%f code=200", r.URL.Path, dur)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/healthz", health)
	http.Handle("/metrics", promhttp.Handler())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Starting on port " + port)
	http.ListenAndServe(":"+port, nil)
}

