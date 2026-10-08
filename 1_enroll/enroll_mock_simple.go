package main

import (
	"net/http"
	"log"
	"encoding/json"
	"time"
	"flag"
)

var (
	delay    = flag.Duration("delay", 20*time.Millisecond, "base response delay")
	slowAt   = flag.Int64("slow-at", 20, "concurrent requests after which the service slows down")
	errorAt  = flag.Int64("error-at", 100, "concurrent requests after which the service may return errors")

	load int64
)

// helpers
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

// ---------------------------------------------------------------------------
func main() {
	http.HandleFunc("/health", healthHandler)

	http.ListenAndServe("127.0.0.1:8080", loggingMiddleware(http.DefaultServeMux))
}

// handlers
func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"status": "ok",
	})	
}
