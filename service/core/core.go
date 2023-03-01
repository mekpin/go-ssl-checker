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
	slackdaily := notification.New("the SSL checker are sucessfully done checking sir! here are the result :") // init slack notification
	slackreminder := notification.New(":rotating_light: *reminder to update the SSL* :rotating_light:")
	// list := make(map[string]model.ExpiryData)
	var (
		// Datapool    []model.ExpiryData
		report      int = 0
		currenttime     = time.Now()
	)

	for _, v := range manifests {
		//debug domain name
		fmt.Println("running check on : " + v.Domainname)

		conn, err := net.Dial("tcp", v.Domainname+":"+v.Domainport)
		if err != nil {
			log.Fatal().Str("message", "error while netdial tcp in file controller/sslcheck.go").Send()
		}

		client := tls.Client(conn, &tls.Config{
			ServerName: v.Domainname,
		})
		defer client.Close()

		if err := client.Handshake(); err != nil {
			log.Fatal().Str("message", "error while client handshake in file controller/sslcheck.go").Send()
		}

		cert := client.ConnectionState().PeerCertificates[0]
		expireddate := cert.NotAfter.Format(time.RFC3339)

		expiredelta := cert.NotAfter.Sub(currenttime)
		deltainteger := int(expiredelta.Hours() / 24)

		//ubah env threshold jadi int untuk compare
		intthreshold, err := strconv.Atoi(config.Common.Threshold)
		if err != nil {
			log.Fatal().Str("message", "error while changing env threshold string to int in file controller/sslcheck.go").Send()
		}

		// list[v.Domainname] = model.ExpiryData{
		// 	Domainname:    v.Domainname,
		// 	Expireddate:   expireddate,
		// 	Remainingdays: deltainteger,
		// }

		list = append(list, model.ExpiryData{
			Domainname:    v.Domainname,
			Expireddate:   expireddate,
			Remainingdays: deltainteger,
		})

		// remind if remaining days are less than the threshold env.

		if deltainteger < intthreshold {
			slackreminder.ReminderSlack(v.Domainname, deltainteger).Send()
			report = report + 1
		}

	}
	slackdaily.SetStatus(nil).ReportCheck(list).Send()

	fmt.Printf("we got %v reports of near expired domain \n", report)

	return list

}
