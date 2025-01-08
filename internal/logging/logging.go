package logging

import (
	"calculate-api/internal/config"
	"fmt"
	"log/slog"
	"os"
)

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

func parseLevel(s string) (slog.Level, error) {
	var lvl slog.Level
	if err := lvl.UnmarshalText([]byte(s)); err != nil {
		return lvl, fmt.Errorf("parseLevel: %w", err)
	}
	return lvl, nil
}
