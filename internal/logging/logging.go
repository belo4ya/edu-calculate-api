package logging

import (
	"calculate-api/internal/config"
	"fmt"
	"log/slog"
	"os"
)

// Configure sets up the logging system based on the provided configuration.
func Configure(conf *config.Config) error {
	lvl, err := parseLevel(conf.LogLevel)
	if err != nil {
		return fmt.Errorf("configure: %w", err)
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     lvl,
		AddSource: true,
	})))
	return nil
}

// parseLevel converts a string representation of the logging level into slog.Level.
func parseLevel(s string) (slog.Level, error) {
	var lvl slog.Level
	if err := lvl.UnmarshalText([]byte(s)); err != nil {
		return lvl, fmt.Errorf("parseLevel: %w", err)
	}
	return lvl, nil
}
