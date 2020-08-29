package config

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	// Server ...
	Server struct {
		Port string `envconfig:"GIN_SERVER_PORT" default:"8088"`
	}

	setting struct {
		Server
	}
)

// Setting config
var Setting = &setting{}

// Load load config
func Load() {
	envconfig.MustProcess("", Setting)
}
