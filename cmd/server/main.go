package main

import (
	"calculate-api/internal/config"
	"calculate-api/internal/logging"
	"calculate-api/internal/server"
	"calculate-api/internal/service"
	"fmt"
	"log"
	"log/slog"

	"github.com/belo4ya/runy"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := runy.SetupSignalHandler()

	conf, err := config.Load(".env")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if err := logging.Configure(conf); err != nil {
		return fmt.Errorf("failed to setup logger: %w", err)
	}

	slog.InfoContext(ctx, "logger is configured")

	slog.InfoContext(ctx, "config initialized", "config", conf.String())

	srv := server.NewHTTPServer(conf)
	svc := service.New()
	svc.RegisterHandlers(srv.Router)

	runy.Add(srv)

	slog.InfoContext(ctx, "starting app")
	if err := runy.Start(ctx); err != nil {
		return fmt.Errorf("problem with running app: %w", err)
	}
	return nil
}
