package middlewares

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/lananolana/webpulse/backend/pkg/http_tools"
)

func Logging(next http.Handler) http.Handler {
	slog.Debug("http logging middleware enabled")

	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		start := time.Now()

		next.ServeHTTP(ww, r)

		slog.Info("http request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.String()),
			slog.String("addr", http_tools.GetRealIP(r)),
			slog.String("user_agent", r.UserAgent()),
			slog.String("id", middleware.GetReqID(r.Context())),
			slog.Int("status", ww.Status()),
			slog.Int("len", ww.BytesWritten()),
			slog.String("duration", time.Since(start).String()),
		)
	}
	return http.HandlerFunc(fn)
}
