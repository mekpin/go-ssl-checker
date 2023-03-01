package main

import (
	"fmt"
	"go-ssl-checker/config"
	"go-ssl-checker/router"
	"go-ssl-checker/service/cron"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	if config.Common.Enablecron == "true" {
		cron.Routine()
		fmt.Println("cron are enabled")
	} else {
		fmt.Println("cron are disabled")
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
