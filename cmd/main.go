package main

import (
	"fmt"
	"go-ssl-checker/config"
	"go-ssl-checker/router"
	"go-ssl-checker/service/cron"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func init() {
	if config.Common.Enablecron == "true" {
		cron.Routine()
		log.Info().Str("message", "cron are enabled").Send()
	} else {
		log.Info().Str("message", "cron are disabled").Send()
	}
}

func main() {

	app := gin.Default()

	router.Router(app)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.Common.Port),
		Handler: app,
	}
	srv.ListenAndServe()

}
