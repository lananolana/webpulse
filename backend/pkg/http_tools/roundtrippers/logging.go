package roundtrippers

import (
	"log/slog"
	"net/http"
	"time"
)

// Logging implements http.RoundTripper.
// Writes logs on each request from http client
type Logging struct {
	next http.RoundTripper
}

func NewLogging() *Logging {
	return &Logging{
		next: http.DefaultTransport,
	}
}

// RoundTrip implements http.RoundTripper.
// Writes logs on each request from http client
func (l *Logging) RoundTrip(r *http.Request) (*http.Response, error) {
	start := time.Now()
	resp, err := l.next.RoundTrip(r)

	slog.Info(
		"httpClient request",
		slog.String("url", r.URL.String()),
		slog.Duration("duration", time.Since(start)),
	)

	return resp, err
}
