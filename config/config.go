package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
)

type common struct {
	Port          string `env:"PORT" envDefault:"3300"`
	Threshold     string `env:"THRESHOLD" envDefault:"30"`
	Croninterval  string `env:"CRON_INTERVAL" envDefault:"daily"`
	Slackwebhook  string `env:"SLACK_WEBHOOK" envDefault:""`
	Enablecron    string `env:"ENABLE_CRON" envDefault:"false"`
	TimeFormat    string `env:"TIME_FORMAT" envDefault:"15:04:05 MST 02 Jan 2006"`
	LocalLocation string `env:"TIME_FORMAT" envDefault:"Asia/Jakarta"`
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
