package config

import "os"

// Config is a config :).
type Config struct {
	HTTPAddr string
}

// Read reads config from environment.
func Read() Config {
	var config Config
	httpAddr, exists := os.LookupEnv("HTTP_ADDR")
	if exists {
		config.HTTPAddr = httpAddr
	}
	return config
}
