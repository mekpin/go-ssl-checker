package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
)

type common struct {
	Port string `env:"PORT" envDefault:"3300"`
}

var (
	Common common
)

func init() {
	env.Parse(&Common)

	log.Debug().Interface("common", Common).Send()
}
