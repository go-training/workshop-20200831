package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Server ...
type Server struct {
	Debug bool   `envconfig:"GIN_DEBUG"`
	Port  string `envconfig:"GIN_SERVER_PORT"`
}

type setting struct {
	Server
	Debug bool `envconfig:"GIN_DEBUG"`
}

// Setting config
var Setting = &setting{}

// Load load config
func Load() {
	envconfig.MustProcess("", Setting)
}
