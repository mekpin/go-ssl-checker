package cron

import (
	"fmt"
	"go-ssl-checker/config"
	"go-ssl-checker/service/core"
	"go-ssl-checker/service/manifest"

	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

func DailyCheck() {
	log.Info().Str("message", "start daily check").Send()
	manifests, err := manifest.ParseInventory(false)
	if err != nil {
		fmt.Println("error while parsing inventory on DailyCheck function in file service/cron/cron.go")
	}

	core.SSLExpireCheck(manifests)

	log.Info().Str("message", "daily check sucessfully ended").Send()

}

func Routine() {
	c := cron.New()
	c.AddFunc("@"+config.Common.Croninterval, DailyCheck)
	// c.AddFunc("@every 10s", DailyCheck)
	//@yearly
	//@monthly
	//@daily
	//@hourly
	c.Start()
}
