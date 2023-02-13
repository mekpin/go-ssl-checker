package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
)

type common struct {
	Port          string `env:"PORT" envDefault:"3300"`
	Domain1       string `env:"DOMAIN1" envDefault:"emptyENV"`
	Domain2       string `env:"DOMAIN2" envDefault:"emptyENV"`
	Domain3       string `env:"DOMAIN3" envDefault:"emptyENV"`
	Domain4       string `env:"DOMAIN4" envDefault:"emptyENV"`
	Domain5       string `env:"DOMAIN5" envDefault:"emptyENV"`
	Alldomainport string `env:"DOMAIN_PORT" envDefault:"443"`
}

var (
	Common common
)

func init() {
	env.Parse(&Common)

	log.Debug().Interface("common", Common).Send()
}
