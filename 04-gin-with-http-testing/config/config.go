package config

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	// Config provides the system configuration.
	Server struct {
		Port string `envconfig:"GIN_SERVER_PORT" default:"8088"`
	}

	setting struct {
		Server
	}
)

var (
	Setting = &setting{}
)

func Load() {
	envconfig.MustProcess("", Setting)
}

// func init() {
// 	envconfig.MustProcess("", Setting)
// }
