package config

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	// Server ...
	Server struct {
		Port string `envconfig:"GIN_SERVER_PORT" default:"8088"`
	}

	// Database ...
	Database struct {
		DBPort string `envconfig:"GIN_DATABASE_PORT" default:"3306"`
	}

	// Cache ...
	Cache struct {
		CachePort string `envconfig:"GIN_CACHE_PORT" default:"6379"`
	}

	setting struct {
		Server
		Database
		Cache
	}
)

// Setting config
var Setting = &setting{}

// Load load config
func Load() {
	envconfig.MustProcess("", Setting)
}
