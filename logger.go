package main

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
)

// Logger wraps an HTTP handler and logs all web requests
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.WithField("method", r.Method).
			WithField("url", r.RequestURI).
			WithField("duration", time.Since(start)).
			Info()
	})
}
