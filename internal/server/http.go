package server

import (
	"calculate-api/internal/config"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
)

type HTTPServer struct {
	HTTP   *http.Server
	Router *chi.Mux
	conf   *config.Config
}

func NewHTTPServer(conf *config.Config) *HTTPServer {
	srv := &HTTPServer{conf: conf}
	srv.Router = chi.NewRouter()
	srv.Router.Use(srv.loggerMiddleware())
	srv.HTTP = &http.Server{Addr: conf.HTTPAddr, Handler: srv.Router}
	return srv
}

func (s *HTTPServer) Start(ctx context.Context) error {
	errCh := make(chan error, 1)
	go func() {
		slog.InfoContext(ctx, fmt.Sprintf("HTTP server start listening on: %s", s.conf.HTTPAddr))
		if err := s.HTTP.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("s.HTTP.ListenAndServe: %w", err)
		}
		close(errCh)
	}()
	select {
	case <-ctx.Done():
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		slog.InfoContext(ctx, "shutting down HTTP server")
		if err := s.HTTP.Shutdown(ctx); err != nil {
			return fmt.Errorf("s.HTTP.Shutdown: %w", err)
		}
		return nil
	case err := <-errCh:
		return err
	}
}

func (s *HTTPServer) loggerMiddleware() func(http.Handler) http.Handler {
	return httplog.Handler(httplog.NewLogger("", httplog.Options{
		LogLevel:       slog.LevelInfo,
		JSON:           true,
		Concise:        true,
		RequestHeaders: true,
	}))
}
