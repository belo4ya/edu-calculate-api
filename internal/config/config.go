package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const EnvLocal = "LOCAL"

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	HTTPAddr    string `mapstructure:"HTTP_ADDR"`
}

func (c *Config) String() string {
	return repr(c,
		"Environment",
		"LogLevel",
		"HTTPAddr",
	)
}

// Load configs override order: defaults -> dotenv -> env.
func Load(dotenvPath string) (*Config, error) {
	conf := Config{
		Environment: EnvLocal,
		LogLevel:    "DEBUG",
		HTTPAddr:    ":8080",
	}

	if _, err := os.Stat(dotenvPath); !errors.Is(err, os.ErrNotExist) {
		if err := godotenv.Load(dotenvPath); err != nil {
			return nil, fmt.Errorf("load: %w", err)
		}
	}

	viper.AutomaticEnv()
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, fmt.Errorf("load: %w", err)
	}

	return &conf, nil
}

func repr(obj any, fields ...string) string {
	v := reflect.ValueOf(obj).Elem()
	parts := make([]string, 0, len(fields))
	for _, field := range fields {
		f := v.FieldByName(field)
		if f.IsValid() {
			parts = append(parts, fmt.Sprintf("%s=%v", field, f.Interface()))
		}
	}
	return fmt.Sprintf("%s{%s}", v.Type().Name(), strings.Join(parts, ", "))
}
