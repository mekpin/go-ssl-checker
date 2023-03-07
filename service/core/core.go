package core

import (
	"crypto/tls"
	"fmt"
	"go-ssl-checker/config"
	"go-ssl-checker/model"
	"go-ssl-checker/service/notification"
	"net"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

func SSLExpireCheck(manifests []model.Inventory) (list []model.ExpiryData) {
	slackdaily := notification.New("the *SSL checker* are done checking my lord! here are the results :") // init slack notification
	slackreminder := notification.New(":rotating_light: *reminder to update the SSL* :rotating_light: \n *ALERT* <!channel>")
	slackerrornotify := notification.New(":warning: *reminder to check the manifest / connectivity*  :warning: \n  <!channel>")

	// list := make(map[string]model.ExpiryData)
	var (
		// Datapool    []model.ExpiryData
		report       int = 0
		errorreport  int = 0
		currenttime      = time.Now()
		reminderlist []model.ExpiryData
		errorlist    []model.ExpiryData
	)

	for _, v := range manifests {
		//debug domain name iteration
		fmt.Println("running check on : " + v.Domainname)

		d := net.Dialer{Timeout: time.Second * 5}
		conn, err := d.Dial("tcp", v.Domainname+":"+v.Domainport)
		if err != nil {
			log.Info().Str("message", "error while netdial tcp in file controller/sslcheck.go skipping loop").Send()
			errorreport = errorreport + 1
		}

		client := tls.Client(conn, &tls.Config{
			ServerName: v.Domainname,
		})
		defer client.Close()

		if err := client.Handshake(); err != nil {
			log.Info().Str("message", "error while client handshake for  in file controller/sslcheck.gob skipping this loop").Send()

			errorreport = errorreport + 1

			list = append(list, model.ExpiryData{
				Domainname:    v.Domainname,
				Expireddate:   "error_found",
				Remainingdays: -50000,
			})

			errorlist = append(reminderlist, model.ExpiryData{
				Domainname:    v.Domainname,
				Expireddate:   "error_found",
				Remainingdays: -50000,
			})

		} else {
			wib, _ := time.LoadLocation(config.Common.LocalLocation)

			cert := client.ConnectionState().PeerCertificates[0]
			expireddate := cert.NotAfter.In(wib).Format(config.Common.TimeFormat)

			expiredelta := cert.NotAfter.Sub(currenttime)
			deltainteger := int(expiredelta.Hours() / 24)

			list = append(list, model.ExpiryData{
				Domainname:    v.Domainname,
				Expireddate:   expireddate,
				Remainingdays: deltainteger,
			})

			// remind if remaining days are less than the threshold env.

			//ubah env threshold jadi int untuk compare
			intthreshold, err := strconv.Atoi(config.Common.Threshold)
			if err != nil {
				log.Info().Str("message", "error while changing env threshold string to int in file controller/sslcheck.go skipping out of loop").Send()
				errorreport = errorreport + 1

			}

			if deltainteger < intthreshold {
				reminderlist = append(reminderlist, model.ExpiryData{
					Domainname:    v.Domainname,
					Expireddate:   expireddate,
					Remainingdays: deltainteger,
				})
				report = report + 1
			}

		}

	}

	slackdaily.ReportCheck(list).Send()
	if reminderlist != nil {
		slackreminder.ReminderSlack(reminderlist).Send()
	}
	if errorlist != nil {
		slackerrornotify.ErrorReportSlack(errorlist).Send()
	}
	fmt.Printf("we got %v reports of near expired domain \n", report)
	fmt.Printf("we got %v error report while SSLcheck \n", errorreport)

	return list

}
