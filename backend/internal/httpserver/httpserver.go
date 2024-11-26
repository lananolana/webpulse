package httpserver

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/lananolana/webpulse/backend/pkg/closer"
	appmiddlewares "github.com/lananolana/webpulse/backend/pkg/http_tools/middlewares"
	_ "github.com/swaggo/files"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Srv represents http server
type Srv struct {
	Server *http.Server
}

// New creates new http server with API routes and middlewares
func New(r *chi.Mux, addr string) *Srv {
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(appmiddlewares.Logging)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.AllowContentType("application/json", "application/text"))
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "OK"}`))
	})

	r.Mount("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./docs"))))
	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/swagger.yaml"),
	))

	return &Srv{
		Server: &http.Server{
			Addr:    addr,
			Handler: r,
		},
	}
}

// MustRun runs http server or panics if error
func (s *Srv) MustRun() {
	go func() {
		if err := s.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic("failed to start http server: " + err.Error())
		}
	}()

	// Graceful shutdown for http server
	closer.Add(func(ctx context.Context) error {
		slog.Debug("shutting down http server")
		return s.Server.Shutdown(ctx)
	})

	slog.Info("http server started on " + s.Server.Addr)
}
