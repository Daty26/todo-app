package core_http_server

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	Addr            string        `envconfig:"ADDR" default:"30s"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" required:"true"`
	AllowedOrigins []string `envconfig:"ALLOWED_ORIGINS"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("HTTP", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}
	return config, nil
}
func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get http server config: %w", err)
		panic(err)
	}
	return config
}
