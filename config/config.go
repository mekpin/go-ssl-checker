package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
)

type common struct {
	Port          string `env:"PORT" envDefault:"3300"`
	Alldomainport string `env:"DOMAIN_PORT" envDefault:"443"`
	Threshold     string `env:"THRESHOLD" envDefault:"25"`
	Croninterval  string `env:"THRESHOLD" envDefault:"daily"`
}

type manifest struct {
	InventoryPath string `env:"MANIFEST_PATH" envDefault:"inventory/manifest.yaml"`
}

var (
	Common   common
	Manifest manifest
)

func init() {
	env.Parse(&Common)
	env.Parse(&Manifest)

	log.Debug().Interface("common", Common).Send()
	log.Debug().Interface("manifest", Manifest).Send()

}
